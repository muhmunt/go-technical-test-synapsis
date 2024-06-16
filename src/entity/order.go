package entity

type Order struct {
	ID          int `gorm:"index"`
	UserID      int
	TotalAmount int
	CreatedAt   int
}

func (Order) TableName() string {
	return "orders"
}
