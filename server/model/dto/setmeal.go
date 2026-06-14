package dto

type SetmealAddDTO struct {
	Name          string              `json:"name" binding:"required"`
	CategoryID    uint64              `json:"categoryId" binding:"required"`
	Price         float64             `json:"price" binding:"required"`
	Image         string              `json:"image" binding:"required"`
	Description   string              `json:"description"`
	SetmealDishes []SetmealDishAddDTO `json:"setmealDishes"`
}

type SetmealDishAddDTO struct {
	DishID uint64 `json:"dishId"`
	Copies int    `json:"copies"`
}

type SetmealUpdateDTO struct {
	ID            uint64              `json:"id" binding:"required"`
	Name          string              `json:"name" binding:"required"`
	CategoryID    uint64              `json:"categoryId" binding:"required"`
	Price         float64             `json:"price" binding:"required"`
	Image         string              `json:"image" binding:"required"`
	Description   string              `json:"description"`
	SetmealDishes []SetmealDishAddDTO `json:"setmealDishes"`
}

type SetmealPageQueryDTO struct {
	Name       string  `form:"name" json:"name"`
	CategoryID *uint64 `form:"categoryId" json:"categoryId"`
	Status     *int    `form:"status" json:"status"`
	Page       int     `form:"page" json:"page"`
	PageSize   int     `form:"pageSize" json:"pageSize"`
}

type SetmealStatusUpdateDTO struct {
	ID     uint64 `form:"id" json:"id"`
	Status int    `uri:"status" json:"status"`
}
