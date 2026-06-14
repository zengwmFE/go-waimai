package entity

type DishFlavor struct {
	ID     uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	DishID uint64 `gorm:"column:dish_id;not null" json:"dishId"`
	Name   string `gorm:"size:32;not null" json:"name"`
	Value  string `gorm:"size:255;not null" json:"value"`
}

func (DishFlavor) TableName() string { return "dish_flavor" }
