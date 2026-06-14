package entity

import "time"

type Orders struct {
	ID                    uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Number                string     `gorm:"size:50;not null;uniqueIndex:uk_number" json:"number"`
	Status                int        `gorm:"not null" json:"status"`
	UserID                uint64     `gorm:"column:user_id;not null" json:"userId"`
	AddressBookID         uint64     `gorm:"column:address_book_id;not null" json:"addressBookId"`
	OrderTime             time.Time  `gorm:"not null" json:"orderTime"`
	CheckoutTime          *time.Time `json:"checkoutTime"`
	PayMethod             *int       `json:"payMethod"`
	PayStatus             int        `gorm:"column:pay_status;not null;default:0" json:"payStatus"`
	Amount                float64    `gorm:"type:decimal(10,2);not null" json:"amount"`
	Remark                string     `gorm:"size:100" json:"remark"`
	Phone                 string     `gorm:"size:11" json:"phone"`
	Address               string     `gorm:"size:255" json:"address"`
	UserName              string     `gorm:"column:user_name;size:32" json:"userName"`
	Consignee             string     `gorm:"size:32" json:"consignee"`
	CancelReason          string     `gorm:"column:cancel_reason;size:255" json:"cancelReason"`
	RejectionReason       string     `gorm:"column:rejection_reason;size:255" json:"rejectionReason"`
	CancelTime            *time.Time `gorm:"column:cancel_time" json:"cancelTime"`
	EstimatedDeliveryTime *time.Time `gorm:"column:estimated_delivery_time" json:"estimatedDeliveryTime"`
	DeliveryStatus        int        `gorm:"column:delivery_status;not null;default:1" json:"deliveryStatus"`
	DeliveryTime          *time.Time `gorm:"column:delivery_time" json:"deliveryTime"`
	PackAmount            int        `gorm:"column:pack_amount;default:0" json:"packAmount"`
	RiderID               *uint64       `gorm:"column:rider_id" json:"riderId"`
	TablewareNumber       int           `gorm:"column:tableware_number;default:0" json:"tablewareNumber"`
	TablewareStatus       int           `gorm:"column:tableware_status;not null;default:1" json:"tablewareStatus"`
	OrderDetails          []OrderDetail `gorm:"foreignKey:OrderID" json:"orderDetails"`
}

func (Orders) TableName() string { return "orders" }
