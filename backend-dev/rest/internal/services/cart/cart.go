package cart

import (
	"fmt"
	"rest/internal/models"
)

func GetCartItemsIDs(cartItems []models.CartItem) ([]int, error) {
	ids := make([]int, len(cartItems))
	for idx, item := range cartItems {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity: %d", item.Quantity)
		}

		ids[idx] = item.ProductID
	}
	return ids, nil
}
