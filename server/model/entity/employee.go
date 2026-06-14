package entity

import "time"

type Employee struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string    `gorm:"size:32;not null" json:"name"`
	Username   string    `gorm:"size:32;not null;uniqueIndex:uk_username" json:"username"`
	Password   string    `gorm:"size:64;not null" json:"password"`
	Phone      string    `gorm:"size:11;not null" json:"phone"`
	Sex        string    `gorm:"size:2;not null" json:"sex"`
	IDNumber   string    `gorm:"column:id_number;size:18;not null" json:"idNumber"`
	Status     int       `gorm:"not null;default:1" json:"status"`
	CreateTime time.Time `gorm:"autoCreateTime" json:"createTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime" json:"updateTime"`
	CreateUser uint64    `gorm:"not null" json:"createUser"`
	UpdateUser uint64    `gorm:"not null" json:"updateUser"`
}

func (Employee) TableName() string { return "employee" }
