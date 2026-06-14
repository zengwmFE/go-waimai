package dao

import (
	"sky-take-out-go/model/dto"
	"sky-take-out-go/model/entity"

	"gorm.io/gorm"
)

type AddressBookDAO struct{ db *gorm.DB }

func NewAddressBookDAO() *AddressBookDAO { return &AddressBookDAO{db: DB} }

// AddressBook
func (dao *AddressBookDAO) GetAddressList(userID uint64) ([]entity.AddressBook, error) {
	var list []entity.AddressBook
	err := dao.db.Where("user_id = ?", userID).Order("is_default DESC, id DESC").Find(&list).Error
	return list, err
}

func (dao *AddressBookDAO) AddAddress(userID uint64, addr dto.AddressBookAddDTO) (uint64, error) {
	book := entity.AddressBook{
		UserID:       userID,
		Consignee:    addr.Consignee,
		Sex:          addr.Sex,
		Phone:        addr.Phone,
		ProvinceCode: addr.ProvinceCode,
		ProvinceName: addr.ProvinceName,
		CityCode:     addr.CityCode,
		CityName:     addr.CityName,
		DistrictCode: addr.DistrictCode,
		DistrictName: addr.DistrictName,
		Detail:       addr.Detail,
		Label:        addr.Label,
	}
	err := dao.db.Create(&book).Error
	return book.ID, err
}

func (dao *AddressBookDAO) GetAddressByID(id, userID uint64) (*entity.AddressBook, error) {
	var addr entity.AddressBook
	err := dao.db.Where("id = ? AND user_id = ?", id, userID).First(&addr).Error
	if err != nil {
		return nil, err
	}
	return &addr, nil
}

func (dao *AddressBookDAO) UpdateAddress(userID uint64, addr dto.AddressBookUpdateDTO) error {
	return dao.db.Model(&entity.AddressBook{}).
		Where("id = ? AND user_id = ?", addr.ID, userID).
		Updates(map[string]any{
			"consignee":     addr.Consignee,
			"sex":           addr.Sex,
			"phone":         addr.Phone,
			"province_code": addr.ProvinceCode,
			"province_name": addr.ProvinceName,
			"city_code":     addr.CityCode,
			"city_name":     addr.CityName,
			"district_code": addr.DistrictCode,
			"district_name": addr.DistrictName,
			"detail":        addr.Detail,
			"label":         addr.Label,
		}).Error
}

func (dao *AddressBookDAO) DeleteAddress(id, userID uint64) error {
	return dao.db.Where("id = ? AND user_id = ?", id, userID).Delete(&entity.AddressBook{}).Error
}

func (dao *AddressBookDAO) SetDefaultAddress(id, userID uint64) error {
	return dao.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&entity.AddressBook{}).Where("user_id = ?", userID).Update("is_default", 0).Error; err != nil {
			return err
		}
		return tx.Model(&entity.AddressBook{}).Where("id = ? AND user_id = ?", id, userID).Update("is_default", 1).Error
	})
}

func (dao *AddressBookDAO) GetDefaultAddress(userID uint64) (*entity.AddressBook, error) {
	var addr entity.AddressBook
	err := dao.db.Where("user_id = ? AND is_default = 1", userID).First(&addr).Error
	if err != nil {
		return nil, err
	}
	return &addr, nil
}
