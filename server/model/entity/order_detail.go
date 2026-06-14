package entity

type OrderDetail struct {
	ID         uint64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string  `gorm:"size:32" json:"name"`
	Image      string  `gorm:"size:255" json:"image"`
	OrderID    uint64  `gorm:"column:order_id;not null" json:"orderId"`
	DishID     *uint64 `gorm:"column:dish_id" json:"dishId"`
	SetmealID  *uint64 `gorm:"column:setmeal_id" json:"setmealId"`
	DishFlavor string  `gorm:"column:dish_flavor;size:50" json:"dishFlavor"`
	Number     int     `gorm:"not null;default:1" json:"number"`
	Amount     float64 `gorm:"type:decimal(10,2);not null" json:"amount"`
}

func (OrderDetail) TableName() string { return "order_detail" }
