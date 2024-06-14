package entity

type Product struct {
	ID          int `gorm:"index"`
	ProductName string
	Price       int
	Description string
	Category    string
	CreatedAt   int
	UpdatedAt   int
}

func (Product) TableName() string {
	return "products"
}
