package dto

type Order struct {
	OrderID    int `json:"order_id"`
	CustomerID int `json:"customer_id"`
	ProductID  int `json:"product_id"`
	Quantity   int `json:"quantity"`
}

type GetOrders struct {
	Orders []*Order `json:"orders"`
}
