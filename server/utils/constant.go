package utils

const (
	CodeSuccess = 1
	CodeError   = 0

	StatusEnabled  = 1
	StatusDisabled = 0

	CategoryTypeDish    = 1
	CategoryTypeSetmeal = 2

	OrderStatusPendingPay    = 1
	OrderStatusPendingAccept = 2
	OrderStatusAccepted      = 3
	OrderStatusDelivering    = 4
	OrderStatusCompleted     = 5
	OrderStatusCancelled     = 6

	PayStatusUnpaid   = 0
	PayStatusPaid     = 1
	PayStatusRefunded = 2

	WsMsgTypeNewOrder = 1
	WsMsgTypeReminder = 2
)
