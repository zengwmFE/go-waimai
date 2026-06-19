package vo

type SetmealDishVO struct {
	DishID uint64  `json:"dishId"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Copies int     `json:"copies"`
}

type SetmealVO struct {
	ID            uint64          `json:"id"`
	Name          string          `json:"name"`
	CategoryID    uint64          `json:"categoryId"`
	Price         float64         `json:"price"`
	Image         string          `json:"image"`
	Description   string          `json:"description"`
	Status        int             `json:"status"`
	SetmealDishes []SetmealDishVO `json:"setmealDishes"`
}

type SetmealOverView struct {
	Sold         int `json:"sold"`
	Discontinued int `json:"discontinued"`
}
