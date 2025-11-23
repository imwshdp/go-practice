package product

import (
	"fmt"
	"net/http"
	"rest/internal/dto"
	"rest/internal/models"
	"rest/internal/storage/postgres/product"
	validateUtil "rest/internal/utils/validator"

	httpUtils "rest/internal/utils/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type productHandler struct {
	productRepo product.ProductRepository
}

func NewProductHandler(productRepo product.ProductRepository) *productHandler {
	return &productHandler{
		productRepo: productRepo,
	}
}

func (h *productHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.getAll).Methods(http.MethodGet)
	router.HandleFunc("/products", h.create).Methods(http.MethodPost)
}

func (h *productHandler) getAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.productRepo.GetAll()
	if err != nil {
		httpUtils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	httpUtils.WriteJSON(w, http.StatusOK, products)
}

func (h *productHandler) create(w http.ResponseWriter, r *http.Request) {
	var product models.CreateProductPayload

	if err := httpUtils.ParseJSON(r, &product); err != nil {
		httpUtils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := validateUtil.Validate.Struct(product); err != nil {
		errors := err.(validator.ValidationErrors)
		httpUtils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	err := h.productRepo.Create(dto.Product{
		Name:        product.Name,
		Description: product.Description,
		Image:       product.Image,
		Price:       product.Price,
		Quantity:    product.Quantity,
	})
	if err != nil {
		httpUtils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	httpUtils.WriteJSON(w, http.StatusCreated, product)
}
