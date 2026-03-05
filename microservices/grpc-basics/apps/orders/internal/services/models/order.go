package models

type Order struct {
	OrderID    int
	CustomerID int
	ProductID  int
	Quantity   int
}

func CreateOrder(
	customerID int,
	productID int,
	quantity int,
) *Order {
	return &Order{
		CustomerID: customerID,
		ProductID:  productID,
		Quantity:   quantity,
	}
}
