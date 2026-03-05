package http

import (
	"grpc-basics/apps/common/http/dto"
	"grpc-basics/apps/common/http/utils"
	"grpc-basics/apps/orders/internal/services"
	"grpc-basics/apps/orders/internal/services/models"
	"net/http"
)

type orderHandler struct {
	ordersService services.OrderService
}

func NewOrderHandler(orderService services.OrderService) *orderHandler {
	return &orderHandler{
		ordersService: orderService,
	}
}

func (h *orderHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *orderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.NewOrder
	err := utils.ParseJSON(r, &req)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	order := models.CreateOrder(
		req.CustomerID,
		req.ProductID,
		req.Quantity,
	)

	err = h.ordersService.CreateOrder(ctx, order)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := &dto.NewOrderResponse{
		Status: dto.StatusSuccess,
	}
	utils.WriteJSON(w, http.StatusOK, res)
}
