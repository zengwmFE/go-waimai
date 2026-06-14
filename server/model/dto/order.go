package dto

import "time"

type OrderPageQueryDTO struct {
	Number    string    `form:"number" json:"number"`
	Phone     string    `form:"phone" json:"phone"`
	Status    int       `form:"status" json:"status"`
	BeginTime time.Time `form:"beginTime" json:"beginTime" time_format:"2006-01-02 15:04:05"`
	EndTime   time.Time `form:"endTime" json:"endTime" time_format:"2006-01-02 15:04:05"`
	Page      int       `form:"page" json:"page"`
	PageSize  int       `form:"pageSize" json:"pageSize"`
}

type OrderConfirmDTO struct {
	ID uint64 `json:"id" binding:"required"`
}

type OrderRejectionDTO struct {
	ID              uint64 `json:"id" binding:"required"`
	RejectionReason string `json:"rejectionReason"`
}

type OrderCancelDTO struct {
	ID           uint64 `json:"id" binding:"required"`
	CancelReason string `json:"cancelReason"`
}

type OrderSubmitDTO struct {
	AddressBookID uint64 `json:"addressBookId" binding:"required"`
	// PayMethod       int    `json:"payMethod" binding:"required"`
	Remark                string  `json:"remark"`
	EstimatedDeliveryTime string  `json:"estimatedDeliveryTime"`
	TablewareNumber       int     `json:"tablewareNumber"`
	TablewareStatus       int     `json:"tablewareStatus"`
	PackAmount            int     `json:"packAmount"`
	Amount                float64 `json:"amount" binding:"required"`
}

type OrderPaymentDTO struct {
	OrderNumber string `json:"orderNumber" binding:"required"`
	PayMethod   int    `json:"payMethod" binding:"required"`
}
