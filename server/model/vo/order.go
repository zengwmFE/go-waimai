package vo

type OrderVO struct{}

type OrderOverView struct {
	WaitingOrders   int `json:"waitingOrders"`
	DeliveredOrders int `json:"deliveredOrders"`
	CompletedOrders int `json:"completedOrders"`
	CancelledOrders int `json:"cancelledOrders"`
	AllOrders       int `json:"allOrders"`
}
