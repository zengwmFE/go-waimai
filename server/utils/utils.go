package utils

import (
	"sky-take-out-go/model/entity"
	"sky-take-out-go/model/vo"
)

func FilterMap(orderList []entity.Orders) vo.OrderOverView {
	var result vo.OrderOverView
	result.AllOrders = len(orderList)
	for _, order := range orderList {
		switch order.Status {
		case OrderStatusPendingPay:
			result.WaitingOrders++
			break
		case OrderStatusPendingAccept:
			result.WaitingOrders++
			break

		case OrderStatusDelivering:
			result.DeliveredOrders++
			break
		case OrderStatusCancelled:
			result.CancelledOrders++
			break
		case OrderStatusCompleted:
			result.CompletedOrders++
			break
		}
	}
	return result
}

func FilterMapDish(dishList []entity.Dish) vo.DishOverView {
	var result vo.DishOverView
	for _, dish := range dishList {
		switch dish.Status {
		case StatusEnabled:
			result.Sold++
			break
		case StatusDisabled:
			result.Discontinued++
			break
		}

	}
	return result
}

func FilterMapSetmeal(setmealList []entity.Setmeal) vo.SetmealOverView {
	var result vo.SetmealOverView
	for _, setmeal := range setmealList {
		switch setmeal.Status {
		case StatusEnabled:
			result.Sold++
			break
		case StatusDisabled:
			result.Discontinued++
			break
		}

	}
	return result
}
