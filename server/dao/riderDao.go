package dao

import (
	"errors"
	"sky-take-out-go/model/entity"
	"sky-take-out-go/utils"

	"gorm.io/gorm"
)

type RiderDAO struct{ db *gorm.DB }

func NewRiderDAO() *RiderDAO { return &RiderDAO{db: DB} }

func (dao *RiderDAO) FindByPhone(phone string) (*entity.Rider, error) {
	var rider entity.Rider
	err := dao.db.Where("phone = ?", phone).First(&rider).Error
	if err != nil {
		return nil, err
	}
	return &rider, nil
}

func (dao *RiderDAO) GetPendingOrders() ([]entity.Orders, error) {
	var orders []entity.Orders
	err := dao.db.Where("status = ?", utils.OrderStatusAccepted).
		Order("order_time ASC").Find(&orders).Error
	return orders, err
}

func (dao *RiderDAO) GrabOrder(orderID, riderID uint64) error {
	result := dao.db.Model(&entity.Orders{}).
		Where("id = ? AND status = ?", orderID, utils.OrderStatusAccepted).
		Updates(map[string]any{
			"status":   utils.OrderStatusDelivering,
			"rider_id": riderID,
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("订单已被其他骑手抢走")
	}
	return nil
}

func (dao *RiderDAO) GetMyOrders(riderID uint64) ([]entity.Orders, error) {
	var orders []entity.Orders
	err := dao.db.Where("status IN ? AND rider_id = ?",
		[]int{utils.OrderStatusDelivering, utils.OrderStatusCompleted}, riderID).
		Order("order_time DESC").Find(&orders).Error
	return orders, err
}

func (dao *RiderDAO) CompleteOrder(orderID, riderID uint64) error {
	result := dao.db.Model(&entity.Orders{}).
		Where("id = ? AND rider_id = ? AND status = ?", orderID, riderID, utils.OrderStatusDelivering).
		Updates(map[string]any{
			"status":        utils.OrderStatusCompleted,
			"delivery_time": gorm.Expr("NOW()"),
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("订单状态异常")
	}
	return nil
}
