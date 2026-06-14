package entity

type AddressBook struct {
	ID           uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID       uint64 `gorm:"column:user_id;not null" json:"userId"`
	Consignee    string `gorm:"size:32;not null" json:"consignee"`
	Sex          string `gorm:"size:2;not null" json:"sex"`
	Phone        string `gorm:"size:11;not null" json:"phone"`
	ProvinceCode string `gorm:"column:province_code;size:12" json:"provinceCode"`
	ProvinceName string `gorm:"column:province_name;size:32" json:"provinceName"`
	CityCode     string `gorm:"column:city_code;size:12" json:"cityCode"`
	CityName     string `gorm:"column:city_name;size:32" json:"cityName"`
	DistrictCode string `gorm:"column:district_code;size:12" json:"districtCode"`
	DistrictName string `gorm:"column:district_name;size:32" json:"districtName"`
	Detail       string `gorm:"size:200" json:"detail"`
	Label        string `gorm:"size:100" json:"label"`
	IsDefault    int    `gorm:"column:is_default;not null;default:0" json:"isDefault"`
}

func (AddressBook) TableName() string { return "address_book" }
