package dto

type UserLoginDTO struct {
	Code string `json:"code" binding:"required"`
}

type ShoppingCartAddDTO struct {
	DishID     *uint64 `json:"dishId"`
	SetmealID  *uint64 `json:"setmealId"`
	DishFlavor string  `json:"dishFlavor"`
}

type ShoppingCartSubDTO struct {
	DishID     *uint64 `json:"dishId"`
	SetmealID  *uint64 `json:"setmealId"`
	DishFlavor string  `json:"dishFlavor"`
}

type AddressBookAddDTO struct {
	Consignee    string `json:"consignee" binding:"required"`
	Sex          string `json:"sex" binding:"required"`
	Phone        string `json:"phone" binding:"required"`
	ProvinceCode string `json:"provinceCode"`
	ProvinceName string `json:"provinceName"`
	CityCode     string `json:"cityCode"`
	CityName     string `json:"cityName"`
	DistrictCode string `json:"districtCode"`
	DistrictName string `json:"districtName"`
	Detail       string `json:"detail"`
	Label        string `json:"label"`
}

type AddressBookUpdateDTO struct {
	ID           uint64 `json:"id" binding:"required"`
	Consignee    string `json:"consignee" binding:"required"`
	Sex          string `json:"sex" binding:"required"`
	Phone        string `json:"phone" binding:"required"`
	ProvinceCode string `json:"provinceCode"`
	ProvinceName string `json:"provinceName"`
	CityCode     string `json:"cityCode"`
	CityName     string `json:"cityName"`
	DistrictCode string `json:"districtCode"`
	DistrictName string `json:"districtName"`
	Detail       string `json:"detail"`
	Label        string `json:"label"`
}

type AddressBookDeleteDTO struct {
	ID uint64 `form:"id" binding:"required"`
}
