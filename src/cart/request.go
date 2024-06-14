package cart

type CartRequest struct {
	ProductID int `json:"product_id" binding:"required"`
	UserID    int `json:"user_id"`
	Quantity  int `json:"quantity" binding:"required"`
}

type CartIDRequest struct {
	ID int `uri:"id" binding:"required"`
}

type CartUserIDRequest struct {
	ID int
}
