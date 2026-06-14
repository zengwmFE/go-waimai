package entity

import "time"

type User struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Openid     string    `gorm:"size:45;not null;uniqueIndex:uk_openid" json:"openid"`
	Name       string    `gorm:"size:32" json:"name"`
	Phone      string    `gorm:"size:11" json:"phone"`
	Sex        string    `gorm:"size:2" json:"sex"`
	IDNumber   string    `gorm:"column:id_number;size:18" json:"idNumber"`
	Avatar     string    `gorm:"size:500" json:"avatar"`
	CreateTime time.Time `gorm:"autoCreateTime" json:"createTime"`
}

func (User) TableName() string { return "user" }
