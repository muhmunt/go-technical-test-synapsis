package entity

type CartItems struct {
	ID        int `gorm:"index"`
	CartID    int
	ProductID int
	Quantity  int
	CreatedAt int
	Product   Product `gorm:"Foreignkey:ProductID;references:ID;"`
}

func (CartItems) TableName() string {
	return "cart_items"
}
