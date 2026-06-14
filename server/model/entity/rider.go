package entity

type Rider struct {
	ID     uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Name   string `gorm:"size:32;not null" json:"name"`
	Phone  string `gorm:"size:11;not null;uniqueIndex" json:"phone"`
	Status int    `gorm:"not null;default:1" json:"status"`
}

func (Rider) TableName() string { return "rider" }
