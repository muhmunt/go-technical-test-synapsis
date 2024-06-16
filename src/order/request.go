package order

type OrderRequest struct {
	UserID      int   `json:"user_id"`
	CartItemsID []int `json:"cart_item" binding:"required"`
}

type OrderIDRequest struct {
	ID string `uri:"id" binding:"required"`
}
