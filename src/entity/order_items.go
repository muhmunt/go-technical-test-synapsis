package entity

type OrderItems struct {
	ID        int `gorm:"index"`
	OrderID   int
	ProductID int
	Quantity  int
	UnitPrice int
	CreatedAt int
	Product   Product `gorm:"Foreignkey:ProductID;references:ID;"`
}

func (OrderItems) TableName() string {
	return "order_items"
}
