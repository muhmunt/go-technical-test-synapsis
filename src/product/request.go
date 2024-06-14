package product

type ProductRequest struct {
	Name        string `json:"name" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Description string `json:"description" binding:"required"`
	Category    string `json:"category" binding:"required"`
}

type ProductIDRequest struct {
	ID int `uri:"id" binding:"required"`
}

type ProductCategoryRequest struct {
	Category string `uri:"category" binding:"required"`
}
