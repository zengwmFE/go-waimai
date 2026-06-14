package dao

import (
	"sky-take-out-go/model/dto"
	"sky-take-out-go/model/entity"

	"gorm.io/gorm"
)

type ShoppingCartDAO struct{ db *gorm.DB }

func NewShoppingCartDAO() *ShoppingCartDAO { return &ShoppingCartDAO{db: DB} }

func (dao *ShoppingCartDAO) AddToCart(userID uint64, req dto.ShoppingCartAddDTO) error {
	// 查菜品或套餐信息
	var name, image string
	var amount float64
	var dishID, setmealID *uint64

	if req.DishID != nil {
		dish, err := NewDishDAO().FindDishById(*req.DishID)
		if err != nil {
			return err
		}
		name = dish.Name
		image = dish.Image
		amount = dish.Price
		dishID = req.DishID
	}
	if req.SetmealID != nil {
		setmeal, _, _, err := NewSetmealDAO().FindSetmealById(*req.SetmealID)
		if err != nil {
			return err
		}
		name = setmeal.Name
		image = setmeal.Image
		amount = setmeal.Price
		setmealID = req.SetmealID
	}

	// 查是否已有同类商品，有则数量+1，没有则插入
	var existing entity.ShoppingCart
	err := dao.db.Where("user_id = ? AND dish_id <=> ? AND setmeal_id <=> ? AND dish_flavor = ?",
		userID, dishID, setmealID, req.DishFlavor).First(&existing).Error
	if err == nil {
		return dao.db.Model(&existing).Update("number", existing.Number+1).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	item := entity.ShoppingCart{
		UserID:     userID,
		DishID:     dishID,
		SetmealID:  setmealID,
		DishFlavor: req.DishFlavor,
		Name:       name,
		Image:      image,
		Amount:     amount,
		Number:     1,
	}
	return dao.db.Create(&item).Error
}

func (dao *ShoppingCartDAO) GetCartList(userID uint64) ([]entity.ShoppingCart, error) {
	var list []entity.ShoppingCart
	err := dao.db.Where("user_id = ?", userID).Order("create_time DESC").Find(&list).Error
	return list, err
}

func (dao *ShoppingCartDAO) SubFromCart(userID uint64, req dto.ShoppingCartSubDTO) error {
	var item entity.ShoppingCart
	err := dao.db.Where("user_id = ? AND dish_id <=> ? AND setmeal_id <=> ? AND dish_flavor = ?",
		userID, req.DishID, req.SetmealID, req.DishFlavor).First(&item).Error
	if err != nil {
		return err
	}
	if item.Number > 1 {
		return dao.db.Model(&item).Update("number", item.Number-1).Error
	}
	return dao.db.Delete(&item).Error
}

func (dao *ShoppingCartDAO) CleanCart(userID uint64) error {
	return dao.db.Where("user_id = ?", userID).Delete(&entity.ShoppingCart{}).Error
}
