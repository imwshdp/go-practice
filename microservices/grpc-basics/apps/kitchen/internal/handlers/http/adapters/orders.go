package adapters

import (
	"grpc-basics/apps/common/genproto/orders"
	"grpc-basics/apps/kitchen/internal/handlers/http/dto"
)

func OrderDtoAdapter(order *orders.Order) *dto.Order {
	return &dto.Order{
		OrderID:    int(order.OrderID),
		CustomerID: int(order.CustomerID),
		ProductID:  int(order.ProductID),
		Quantity:   int(order.Quantity),
	}
}

func OrdersDtoAdapter(data []*orders.Order) []*dto.Order {
	dtoOrders := make([]*dto.Order, 0, len(data))

	for _, or := range data {
		dtoOrders = append(dtoOrders, OrderDtoAdapter(or))
	}

	return dtoOrders
}
