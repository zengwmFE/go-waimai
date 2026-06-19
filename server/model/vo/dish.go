package vo

type DishVO struct{}

type DishOverView struct {
	Sold         int `json:"sold"`
	Discontinued int `json:"discontinued"`
}
