package entity

type Cart struct {
	ID        int `gorm:"index"`
	UserID    int
	CreatedAt int
	CartItems []CartItems `gorm:"foreignKey:CartID"`
	User      *User       `gorm:"Foreignkey:user_id;references:id;"`
}

func (Cart) TableName() string {
	return "carts"
}
