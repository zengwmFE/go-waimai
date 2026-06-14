package entity

import "time"

type Setmeal struct {
	ID            uint64        `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string        `gorm:"size:32;not null;uniqueIndex:uk_name" json:"name"`
	CategoryID    uint64        `gorm:"column:category_id;not null" json:"categoryId"`
	Price         float64       `gorm:"type:decimal(10,2);not null" json:"price"`
	Image         string        `gorm:"size:255;not null" json:"image"`
	Description   string        `gorm:"size:255" json:"description"`
	Status        int           `gorm:"not null;default:1" json:"status"`
	CreateTime    time.Time     `gorm:"autoCreateTime" json:"createTime"`
	UpdateTime    time.Time     `gorm:"autoUpdateTime" json:"updateTime"`
	CreateUser    uint64        `gorm:"not null" json:"createUser"`
	UpdateUser    uint64        `gorm:"not null" json:"updateUser"`
	SetmealDishes []SetmealDish `gorm:"foreignKey:SetmealID" json:"setmealDishes"`
}

func (Setmeal) TableName() string { return "setmeal" }
