package cart

import (
	"go-technical-test-synapsis/src/entity"
)

type CartFormatter struct {
	ID        int         `json:"id"`
	UserID    int         `json:"user_id"`
	CartItems interface{} `json:"cart_items"`
	CreatedAt int         `json:"created_at"`
}

func FormatCart(cart entity.Cart) CartFormatter {
	formatter := CartFormatter{
		ID:        cart.ID,
		UserID:    cart.UserID,
		CreatedAt: cart.CreatedAt,
		CartItems: cart.CartItems,
	}

	return formatter
}

func FormatCarts(carts []entity.Cart) []CartFormatter {
	cartsFormatter := []CartFormatter{}

	for _, cart := range carts {
		cartFormatter := FormatCart(cart)
		cartsFormatter = append(cartsFormatter, cartFormatter)
	}

	return cartsFormatter
}
