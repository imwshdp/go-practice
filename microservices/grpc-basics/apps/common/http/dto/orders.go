package dto

type status = string

const (
	StatusSuccess = "success"
)

type NewOrder struct {
	CustomerID int `json:"customer_id"`
	ProductID  int `json:"product_id"`
	Quantity   int `json:"quantity"`
}

type NewOrderResponse struct {
	Status status `json:"status"`
}
