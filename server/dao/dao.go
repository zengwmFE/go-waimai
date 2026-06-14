package dao

import (
	"crypto/rand"
	"errors"
	"fmt"
	"sky-take-out-go/model/dto"
	"sky-take-out-go/model/entity"
	"sky-take-out-go/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var DB *gorm.DB

type EmployeeDAO struct{ db *gorm.DB }
type CategoryDAO struct{ db *gorm.DB }
type DishDAO struct{ db *gorm.DB }
type SetmealDAO struct{ db *gorm.DB }
type OrderDAO struct{ db *gorm.DB }
type UserDAO struct{ db *gorm.DB }

func NewEmployeeDAO() *EmployeeDAO { return &EmployeeDAO{db: DB} }
func NewCategoryDAO() *CategoryDAO { return &CategoryDAO{db: DB} }
func NewDishDAO() *DishDAO         { return &DishDAO{db: DB} }
func NewSetmealDAO() *SetmealDAO   { return &SetmealDAO{db: DB} }
func NewOrderDAO() *OrderDAO       { return &OrderDAO{db: DB} }
func NewUserDAO() *UserDAO         { return &UserDAO{db: DB} }

// 根据id查找User表是否存在

func (dao *UserDAO) FindByUserId(openid string) (*entity.User, error) {
	var user entity.User
	result := dao.db.Where("openid = ?", openid).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func (dao *UserDAO) CreateUserId(openid string) (*entity.User, error) {
	user := entity.User{
		Openid: openid,
	}
	result := dao.db.Create(&user)
	if result.Error != nil {
		return &user, nil
	}

	return nil, result.Error
}

/*
*
Admin
*/
func (dao *EmployeeDAO) FindByUserName(username string) (*entity.Employee, error) {
	employee := entity.Employee{
		Username: username,
	}

	result := dao.db.Where("username = ?", username).First(&employee)
	if result.Error != nil {
		return nil, result.Error
	}

	return &employee, nil
}

func (dao *EmployeeDAO) FindByUserId(id uint64) (*entity.Employee, error) {
	employee := entity.Employee{
		ID: id,
	}

	result := dao.db.Where("id = ?", id).First(&employee)
	if result.Error != nil {
		return nil, result.Error
	}

	return &employee, nil
}

func (dao *EmployeeDAO) CreateUser(username string, password string) (*entity.Employee, error) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	employee := entity.Employee{
		Username: username,
		Password: string(hashed),
	}
	result := dao.db.Create(&employee)
	if result.Error != nil {
		return nil, result.Error
	}
	return &employee, nil
}

func (dao *EmployeeDAO) CreateUserByModel(dto *dto.EmployeeAddDTO) (*entity.Employee, error) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	employee := entity.Employee{
		Name:     dto.Name,
		Username: dto.Username,
		Phone:    dto.Phone,
		Sex:      dto.Sex,
		IDNumber: dto.IDNumber,
		Password: string(hashed),
		Status:   1, // 默认启用
	}
	result := dao.db.Create(&employee)
	if result.Error != nil {
		return nil, result.Error
	}
	return &employee, nil
}

func (dao *EmployeeDAO) FindAllUser(dto *dto.EmployeePageQueryDTO) ([]entity.Employee, int64, error) {
	var employee []entity.Employee
	var total int64
	query := dao.db.Where(`name like ?`, "%"+dto.Name+"%")

	result := query.Limit(dto.PageSize).Offset((dto.Page - 1) * dto.PageSize).Find(&employee)
	query.Count(&total)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return employee, total, result.Error
}

func (dao *EmployeeDAO) EditUser(dto *dto.EmployeeUpdateDTO) error {
	employee := entity.Employee{
		ID:       dto.ID,
		IDNumber: dto.IDNumber,
		Sex:      dto.Sex,
		Phone:    dto.Phone,
		Name:     dto.Name,
		Username: dto.Username,
	}
	result := dao.db.Updates(&employee)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (dao *EmployeeDAO) UpdateUserStatusById(id uint64, status int) (uint64, error) {
	result := dao.db.Model(&entity.Employee{}).Where("id=?", id).Update("status", status)
	if result.Error != nil {

		return id, result.Error
	}
	return id, nil
}

func (dao *EmployeeDAO) UpdatePwdByEmpId(empId uint64, newPassword string) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	result := dao.db.Model(&entity.Employee{}).Where("id = ?", empId).Update("password", hashed)
	if result.Error != nil {

		return result.Error
	}
	return nil
}

// Category
func (dao *CategoryDAO) FindCategoryByName(catgoryName string) (*entity.Category, error) {
	var category = entity.Category{}
	result := dao.db.Where(`name=?`, catgoryName).First(&category)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}
	return &category, nil
}

func (dao *CategoryDAO) InsertCategory(category *dto.CategoryAddDTO) (uint64, error) {
	categoryAdd := entity.Category{
		Name: category.Name,
		Sort: category.Sort,
		Type: category.Type,
	}
	result := dao.db.Create(&categoryAdd)
	if result.Error != nil {
		return 0, result.Error
	}
	return categoryAdd.ID, nil
}

func (dao *CategoryDAO) FindAllCategory(searchQuery *dto.CategoryPageQueryDTO) ([]entity.Category, int64, error) {
	limit := searchQuery.PageSize
	offset := (searchQuery.Page - 1) * searchQuery.PageSize
	name := searchQuery.Name
	Type := searchQuery.Type
	fmt.Print(limit, offset, name, Type)
	var categoryList []entity.Category
	var total int64
	query := dao.db.Where(`name LIKE ?`, "%"+name+"%")
	if Type != nil {
		query = query.Where(`type =?`, Type)
	}
	result := query.Limit(limit).Offset(offset).Find(&categoryList)
	query.Count(&total)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return categoryList, total, nil
}

func (dao *CategoryDAO) DeleteCategoryById(id uint64) error {
	var category entity.Category
	result := dao.db.Where("id=?", id).Delete(&category)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (dao *CategoryDAO) UpdateCategoryById(id uint64, status uint64) error {
	result := dao.db.Model(&entity.Category{}).Where(`id=?`, id).Update("status", status)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (dao *CategoryDAO) EditCategoryById(id uint64, name string, sort int) error {
	result := dao.db.Model(&entity.Category{}).Where(`id=?`, id).Updates((map[string]interface{}{"name": name,
		"sort": sort}))
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (dao *CategoryDAO) FindAllCategoryByType(categoryType int) (*[]entity.Category, error) {
	var categoryList []entity.Category
	result := dao.db.Where(`type=?`, categoryType).Find(&categoryList)
	if result.Error != nil {
		return nil, result.Error
	}
	return &categoryList, nil
}

// dish
func (dao *DishDAO) FindAllDish(dishQuery *dto.DishPageQueryDTO) (*[]entity.Dish, int64, error) {
	var dish []entity.Dish
	var total int64
	db := dao.db.Model(&dish)
	if dishQuery.Status != nil {
		db = db.Where(`status=?`, *dishQuery.Status)
	}
	if dishQuery.CategoryID != nil {
		db = db.Where(`category_id=?`, *dishQuery.CategoryID)
	}
	if dishQuery.Name != "" {
		db = db.Where("name LIKE ?", "%"+dishQuery.Name+"%")
	}
	db.Count(&total)
	db = db.Limit(dishQuery.PageSize).Offset((dishQuery.Page - 1) * dishQuery.PageSize)
	result := db.Find(&dish)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return &dish, total, nil
}

func (dao *DishDAO) AddDish(dishJSON *dto.DishAddDTO) (*entity.Dish, error) {
	flavors := make([]entity.DishFlavor, len(dishJSON.Flavors))
	for i, f := range dishJSON.Flavors {
		flavors[i] = entity.DishFlavor{Name: f.Name, Value: f.Value}
	}
	var dish = entity.Dish{
		Name:        dishJSON.Name,
		CategoryID:  dishJSON.CategoryID,
		Price:       dishJSON.Price,
		Image:       dishJSON.Image,
		Description: dishJSON.Description,
		Flavors:     flavors,
	}
	result := dao.db.Create(&dish)
	if result.Error != nil {
		return nil, result.Error
	}

	return &dish, nil
}

func (dao *DishDAO) UpdateDish(dishJSON *dto.DishUpdateDTO) (*entity.Dish, error) {
	flavors := make([]entity.DishFlavor, len(dishJSON.Flavors))
	for i, f := range dishJSON.Flavors {
		flavors[i] = entity.DishFlavor{
			Name:  f.Name,
			Value: f.Value,
		}
	}
	var dish = entity.Dish{
		ID:          dishJSON.ID,
		Name:        dishJSON.Name,
		CategoryID:  dishJSON.CategoryID,
		Price:       dishJSON.Price,
		Image:       dishJSON.Image,
		Flavors:     flavors,
		Description: dishJSON.Description,
	}
	result := dao.db.Model(&dish).Updates(&dish)
	if result.Error != nil {
		return nil, result.Error
	}

	return &dish, nil
}

func (dao *DishDAO) FindDishById(id uint64) (*entity.Dish, error) {
	var dish = entity.Dish{
		ID: id,
	}
	result := dao.db.Preload("Flavors").First(&dish)

	if result.Error != nil {
		return nil, result.Error
	}
	return &dish, nil
}

func (dao *DishDAO) DeleteDishByIds(ids []uint64) error {
	var dishs = make([]entity.Dish, len(ids))
	for index, id := range ids {
		dishs[index] = entity.Dish{
			ID: id,
		}
	}
	result := dao.db.Delete(&dishs)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (dao *DishDAO) UpdateStatusById(status int, id uint64) error {
	var dishResult entity.Dish
	result := dao.db.Model(&dishResult).Where(`id=?`, id).Update("status", status)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// setmeal
func (dao *SetmealDAO) AddSetmeal(setmeal *dto.SetmealAddDTO) (*entity.Setmeal, error) {
	setmealDish := make([]entity.SetmealDish, len(setmeal.SetmealDishes))

	for index, dish := range setmeal.SetmealDishes {
		setmealDish[index] = entity.SetmealDish{
			DishID: dish.DishID,
			Copies: dish.Copies,
		}
	}

	var setmealEntity = entity.Setmeal{
		Name:          setmeal.Name,
		CategoryID:    setmeal.CategoryID,
		Price:         setmeal.Price,
		Image:         setmeal.Image,
		Description:   setmeal.Description,
		SetmealDishes: setmealDish,
	}

	result := dao.db.Create(&setmealEntity)
	if result.Error != nil {

		return nil, result.Error
	}
	return &setmealEntity, nil
}

// find dish by categoryId
func (dao *DishDAO) FindDishBycategoryId(categoryId uint64) ([]entity.Dish, error) {
	var dishEntity []entity.Dish
	result := dao.db.Preload("Flavors").Where(`category_id=?`, categoryId).Find(&dishEntity)
	if result.Error != nil {
		return nil, result.Error
	}

	return dishEntity, nil
}

func (dao *SetmealDAO) FindAllSetmeal(setmeal *dto.SetmealPageQueryDTO) ([]entity.Setmeal, int64, error) {
	var setsealEntity []entity.Setmeal
	var total int64
	db := dao.db.Model(&setsealEntity).Preload("SetmealDishes")
	if setmeal.Name != "" {
		db = db.Where(`name LIKE`, "%"+setmeal.Name+"%")
	}

	if setmeal.CategoryID != nil {
		db = db.Where(`category_id=?`, setmeal.CategoryID)
	}
	if setmeal.Status != nil {
		db = db.Where(`status=?`, *setmeal.Status)
	}
	db = db.Count(&total)
	db = db.Limit(setmeal.PageSize).Offset((setmeal.Page - 1) * setmeal.PageSize)
	result := db.Find(&setsealEntity)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return setsealEntity, total, nil
}

func (dao *SetmealDAO) DeleteSetmealByIds(ids []uint64) error {
	var setmeal = make([]entity.Setmeal, len(ids))
	for index, id := range ids {
		setmeal[index] = entity.Setmeal{
			ID: id,
		}
	}
	result := dao.db.Delete(&setmeal)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (dao *SetmealDAO) UpdateSetmeal(setmealJSON *dto.SetmealUpdateDTO) (*entity.Setmeal, error) {

	setmealDishes := make([]entity.SetmealDish, len(setmealJSON.SetmealDishes))
	for i, f := range setmealJSON.SetmealDishes {
		setmealDishes[i] = entity.SetmealDish{
			Copies: f.Copies,
			DishID: f.DishID,
		}
	}
	var setmeal = entity.Setmeal{
		ID:            setmealJSON.ID,
		Name:          setmealJSON.Name,
		CategoryID:    setmealJSON.CategoryID,
		Price:         setmealJSON.Price,
		Image:         setmealJSON.Image,
		SetmealDishes: setmealDishes,
		Description:   setmealJSON.Description,
	}
	// 事务来保证一系列的数据
	err := dao.db.Transaction(func(tx *gorm.DB) error {
		// 先整体删除
		if err := tx.Where(`setmeal_id=?`, setmealJSON.ID).Delete(&setmealDishes).Error; err != nil {
			return err
		}
		// 数据插入
		if err := tx.Create(&setmealDishes).Error; err != nil {
			return err
		}
		// 更新主表数据
		if err := tx.Model(&setmeal).Updates(&setmeal).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &setmeal, nil
}

func (dao *SetmealDAO) FindSetmealById(id uint64) (*entity.Setmeal, []entity.SetmealDish, []entity.Dish, error) {
	var setmeal entity.Setmeal
	if err := dao.db.First(&setmeal, id).Error; err != nil {
		return nil, nil, nil, err
	}

	var sd []entity.SetmealDish
	if err := dao.db.Where("setmeal_id = ?", id).Find(&sd).Error; err != nil {
		return nil, nil, nil, err
	}

	// 收集所有 dishId，批量查
	dishIDs := make([]uint64, len(sd))
	for i, v := range sd {
		dishIDs[i] = v.DishID
	}

	var dishes []entity.Dish
	if len(dishIDs) > 0 {
		if err := dao.db.Where("id IN ?", dishIDs).Find(&dishes).Error; err != nil {
			return nil, nil, nil, err
		}
	}

	return &setmeal, sd, dishes, nil
}

// // Order
func (dao *OrderDAO) SubmitOrder(userId uint64, orderSubmit dto.OrderSubmitDTO) (*entity.Orders, error) {
	now := time.Now()
	deliveryTime, _ := time.Parse("2006-01-02 15:04:05", orderSubmit.EstimatedDeliveryTime)

	// 时间戳 + 真随机数防碰撞
	b := make([]byte, 3)
	rand.Read(b)
	orderNumber := fmt.Sprintf("%s%06d", now.Format("20060102150405"), (int(b[0])<<8|int(b[1]))%1000000)

	// 查询地址信息，写入订单快照
	addr, err := NewAddressBookDAO().GetAddressByID(orderSubmit.AddressBookID, userId)
	if err != nil {
		return nil, err
	}

	var orderEntity = entity.Orders{
		Number:                orderNumber,
		Status:                utils.OrderStatusPendingPay,
		UserID:                userId,
		AddressBookID:         orderSubmit.AddressBookID,
		Amount:                orderSubmit.Amount,
		Remark:                orderSubmit.Remark,
		PackAmount:            orderSubmit.PackAmount,
		TablewareNumber:       orderSubmit.TablewareNumber,
		TablewareStatus:       orderSubmit.TablewareStatus,
		EstimatedDeliveryTime: &deliveryTime,
		OrderTime:             now,
		PayStatus:             utils.PayStatusUnpaid,
		Phone:                 addr.Phone,
		Consignee:             addr.Consignee,
		Address:               addr.ProvinceName + addr.CityName + addr.DistrictName + addr.Detail,
	}

	err = dao.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&orderEntity).Error; err != nil {
			return err
		}
		if err := NewShoppingCartDAO().CleanCart(userId); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &orderEntity, nil
}

func (dao *OrderDAO) ConditionSearch(query *dto.OrderPageQueryDTO) ([]entity.Orders, int64, error) {
	var list []entity.Orders
	var total int64
	db := dao.db.Model(&entity.Orders{})
	if query.Number != "" {
		db = db.Where("number LIKE ?", "%"+query.Number+"%")
	}
	if query.Phone != "" {
		db = db.Where("phone LIKE ?", "%"+query.Phone+"%")
	}
	if query.Status != 0 {
		db = db.Where("status = ?", query.Status)
	}
	if !query.BeginTime.IsZero() {
		db = db.Where("order_time >= ?", query.BeginTime)
	}
	if !query.EndTime.IsZero() {
		db = db.Where("order_time <= ?", query.EndTime)
	}
	db.Count(&total)
	result := db.Order("order_time DESC").Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&list)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return list, total, nil
}

func (dao *OrderDAO) GetOrderByID(id uint64) (*entity.Orders, error) {
	var order entity.Orders
	err := dao.db.Preload("OrderDetails").First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

type OrderStatisticsVO struct {
	Status int   `json:"status"`
	Count  int64 `json:"count"`
}

func (dao *OrderDAO) Statistics() ([]OrderStatisticsVO, error) {
	var result []OrderStatisticsVO
	err := dao.db.Model(&entity.Orders{}).
		Select("status, COUNT(*) AS count").
		Group("status").
		Order("status").
		Scan(&result).Error
	return result, err
}

func (dao *OrderDAO) PayOrder(userID uint64, orderNumber string, payMethod int) error {
	result := dao.db.Model(&entity.Orders{}).
		Where("number = ? AND user_id = ? AND status = ?", orderNumber, userID, utils.OrderStatusPendingPay).
		Updates(map[string]any{
			"status":        utils.OrderStatusPendingAccept,
			"pay_status":    utils.PayStatusPaid,
			"pay_method":    payMethod,
			"checkout_time": time.Now(),
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("订单不存在或已支付")
	}
	return nil
}

// confirm_order

func (dao *OrderDAO) ConfirmOrder(id uint64) (uint64, error) {
	result := dao.db.Model(&entity.Orders{}).
		Where(`id=?`, id).Updates(map[string]any{
		"status": utils.OrderStatusAccepted,
	})
	if result.Error != nil {
		return id, result.Error
	}

	return id, nil
}
