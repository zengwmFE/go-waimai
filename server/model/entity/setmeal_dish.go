package entity

type SetmealDish struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	SetmealID uint64 `gorm:"column:setmeal_id;not null" json:"setmealId"`
	DishID    uint64 `gorm:"column:dish_id;not null" json:"dishId"`
	Copies    int    `gorm:"not null;default:1" json:"copies"`
}

func (SetmealDish) TableName() string { return "setmeal_dish" }
