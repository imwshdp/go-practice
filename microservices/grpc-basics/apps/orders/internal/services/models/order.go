package models

type Order struct {
	OrderID    int
	CustomerID int
	ProductID  int
	Quantity   int
}

func CreateOrder(
	orderID int,
	customerID int,
	productID int,
	quantity int,
) *Order {
	return &Order{
		OrderID:    orderID,
		CustomerID: customerID,
		ProductID:  productID,
		Quantity:   quantity,
	}
}
