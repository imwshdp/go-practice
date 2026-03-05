package http

import (
	"context"
	"grpc-basics/apps/common/genproto/orders"
	commonDto "grpc-basics/apps/common/http/dto"
	"grpc-basics/apps/common/http/utils"
	"grpc-basics/apps/kitchen/internal/handlers/http/adapters"
	"grpc-basics/apps/kitchen/internal/handlers/http/dto"
	"grpc-basics/apps/kitchen/internal/web/templates"
	"html/template"
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
	router.HandleFunc("GET /orders", h.GetOrders)
	router.HandleFunc("GET /report", h.GetOrdersReport)
}

func (h *orderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	var req commonDto.NewOrder
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

	res := &commonDto.NewOrderResponse{
		Status: commonDto.StatusSuccess,
	}
	utils.WriteJSON(w, http.StatusOK, res)
}

func (h *orderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	data, err := h.grpcClient.GetOrders(ctx, &orders.GetOrdersRequest{})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	dtoOrders := adapters.OrdersDtoAdapter(data.Orders)

	res := &dto.GetOrders{
		Orders: dtoOrders,
	}
	utils.WriteJSON(w, http.StatusOK, res)
}

func (h *orderHandler) GetOrdersReport(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	data, err := h.grpcClient.GetOrders(ctx, &orders.GetOrdersRequest{})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	dtoOrders := adapters.OrdersDtoAdapter(data.Orders)

	ordersTemplate := template.Must(template.New("orders").Parse(templates.OrdersTemplate))

	if err = ordersTemplate.Execute(w, dtoOrders); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}
