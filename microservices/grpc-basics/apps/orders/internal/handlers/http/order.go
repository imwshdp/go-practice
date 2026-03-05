package http

import (
	httpUtils "grpc-basics/apps/common/utils/http"
	"grpc-basics/apps/orders/internal/handlers/http/dto"
	"grpc-basics/apps/orders/internal/services"
	"grpc-basics/apps/orders/internal/services/models"
	"net/http"
)

type OrdersHandler struct {
	ordersService services.OrderService
}

func NewOrderHandler(orderService services.OrderService) *OrdersHandler {
	return &OrdersHandler{
		ordersService: orderService,
	}
}

func (h *OrdersHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrdersHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.NewOrder
	err := httpUtils.ParseJSON(r, &req)
	if err != nil {
		httpUtils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	order := models.CreateOrder(
		42,
		req.CustomerID,
		req.ProductID,
		req.Quantity,
	)

	err = h.ordersService.CreateOrder(ctx, order)
	if err != nil {
		httpUtils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := &dto.NewOrderResponse{
		Status: dto.StatusSuccess,
	}
	httpUtils.WriteJSON(w, http.StatusOK, res)
}
