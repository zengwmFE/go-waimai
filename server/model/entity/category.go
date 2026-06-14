package entity

import "time"

type Category struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string    `gorm:"size:32;not null;uniqueIndex:uk_name" json:"name"`
	Type       int       `gorm:"not null" json:"type"`
	Sort       int       `gorm:"not null;default:0" json:"sort"`
	Status     int       `gorm:"not null;default:1" json:"status"`
	CreateTime time.Time `gorm:"autoCreateTime" json:"createTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime" json:"updateTime"`
	CreateUser uint64    `gorm:"not null" json:"createUser"`
	UpdateUser uint64    `gorm:"not null" json:"updateUser"`
}

func (Category) TableName() string { return "category" }
