package http

import (
	"context"
	"grpc-basics/apps/common/genproto/orders"
	"grpc-basics/apps/common/http/dto"
	"grpc-basics/apps/common/http/utils"
	"net/http"
	"time"
)

type orderHandler struct {
	grpcClient orders.OrderServiceClient
}

func NewOrderHandler(
	grpcClient orders.OrderServiceClient,
) *orderHandler {
	return &orderHandler{
		grpcClient,
	}
}

func (h *orderHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *orderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	var req dto.NewOrder
	err := utils.ParseJSON(r, &req)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	_, err = h.grpcClient.CreateOrder(ctx, &orders.CreateOrderRequest{
		CustomerID: int32(req.CustomerID),
		ProductID:  int32(req.ProductID),
		Quantity:   int32(req.Quantity),
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := &dto.NewOrderResponse{
		Status: dto.StatusSuccess,
	}
	utils.WriteJSON(w, http.StatusOK, res)
}
