package cart

import (
	"fmt"
	"net/http"
	"rest/internal/dto"
	"rest/internal/models"
	"rest/internal/services/auth"
	cartService "rest/internal/services/cart"
	"rest/internal/storage/postgres/order"
	"rest/internal/storage/postgres/product"
	"rest/internal/storage/postgres/user"
	httpUtils "rest/internal/utils/http"
	validateUtil "rest/internal/utils/validator"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type cartHandler struct {
	orderRepo   order.OrderRepository
	productRepo product.ProductRepository
	userRepo    user.UserRepository
}

func NewOrderHandler(cartRepo order.OrderRepository, productRepo product.ProductRepository, userRepo user.UserRepository) *cartHandler {
	return &cartHandler{
		orderRepo:   cartRepo,
		productRepo: productRepo,
		userRepo:    userRepo,
	}
}

func (h *cartHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", auth.WithJWTAuth(h.checkout, h.userRepo)).Methods(http.MethodPost)
}

func (h *cartHandler) checkout(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserFromCtx(r).ID

	var cart models.CartCheckoutPayload

	if err := httpUtils.ParseJSON(r, &cart); err != nil {
		httpUtils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := validateUtil.Validate.Struct(cart); err != nil {
		errors := err.(validator.ValidationErrors)
		httpUtils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid payload %v", errors))
		return
	}

	productIDs, err := cartService.GetCartItemsIDs(cart.Items)
	if err != nil {
		httpUtils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	products, err := h.productRepo.GetByIds(productIDs)
	if err != nil {
		httpUtils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	orderID, totalPrice, err := h.createOrder(products, cart.Items, userID)
	if err != nil {
		httpUtils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	httpUtils.WriteJSON(w, http.StatusCreated, map[string]any{
		"order_id":    orderID,
		"total_price": totalPrice,
	})
}

func (h *cartHandler) createOrder(products []*dto.Product, items []models.CartItem, userID int) (int, float64, error) {
	productsMap := make(map[int]*dto.Product)

	for _, product := range products {
		productsMap[product.ID] = product
	}

	if err := checkIfCartIsInStock(items, productsMap); err != nil {
		return 0, 0, err
	}

	totalPrice := calculateTotalPrice(items, productsMap)

	for _, item := range items {
		product := productsMap[item.ProductID]
		product.Quantity -= item.Quantity

		h.productRepo.Update(product)
	}

	orderID, err := h.orderRepo.Create(dto.Order{
		UserID:  userID,
		Total:   totalPrice,
		Status:  "pending",
		Address: "TODO address",
	})

	if err != nil {
		return 0, 0, fmt.Errorf("order creating", err)
	}

	for _, item := range items {
		product := productsMap[item.ProductID]
		err := h.orderRepo.CreateOrderItem(dto.OrderItem{
			OrderID:   orderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		})

		if err != nil {
			return 0, 0, fmt.Errorf("order item creation: %s", err)
		}
	}

	return orderID, totalPrice, nil
}

func calculateTotalPrice(items []models.CartItem, productsMap map[int]*dto.Product) float64 {
	totalPrice := 0.0

	for _, item := range items {
		product := productsMap[item.ProductID]
		totalPrice += product.Price * float64(item.Quantity)
	}

	return totalPrice
}

func checkIfCartIsInStock(items []models.CartItem, productsMap map[int]*dto.Product) error {
	if len(items) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range items {
		product, ok := productsMap[item.ProductID]
		if !ok {
			return fmt.Errorf("product %d is not presented", item.ProductID)
		}

		if product.Quantity < item.Quantity {
			return fmt.Errorf("product %d is out of stock", item.ProductID)
		}
	}

	return nil
}
