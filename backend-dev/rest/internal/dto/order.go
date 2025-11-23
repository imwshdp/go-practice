package dto

type Order struct {
	ID        int
	UserID    int
	Total     float64
	Status    string
	Address   string
	CreatedAt string
}
