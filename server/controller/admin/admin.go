package admin

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sky-take-out-go/config"
	"sky-take-out-go/dao"
	"sky-take-out-go/model/dto"
	"sky-take-out-go/model/entity"
	"sky-take-out-go/model/vo"
	"sky-take-out-go/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Employee
func Login(c *gin.Context) {
	var request dto.EmployeeLoginDTO
	if error := c.ShouldBindJSON(&request); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": error.Error(),
		})
		return
	}

	username := request.Username
	password := request.Password
	employeeDao := dao.NewEmployeeDAO()
	employee, _ := employeeDao.FindByUserName(username)
	if employee == nil {
		// init的时候，数据库没有admin，增加了，后面去掉，添加上正常判断
		result, error := employeeDao.CreateUser(username, password)
		if error != nil {
			c.JSON(http.StatusBadRequest, vo.EmployeeLoginError{
				Code:   400,
				ErrMsg: "错误",
			})
			return
		}
		token, err := utils.GenerateAdminToken(config.Cfg.JWT.AdminSecret, config.Cfg.JWT.AdminTTL, result.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, vo.EmployeeLoginError{
				Code:   400,
				ErrMsg: "错误",
			})
			return
		}
		c.JSON(http.StatusOK, vo.EmployeeLoginResponseVO{
			Code: 1,
			Data: vo.EmployeeLoginVO{
				ID:       result.ID,
				Name:     result.Name,
				Username: result.Name,
				Token:    token,
			},
		})
		// c.JSON(http.StatusOK, vo.EmployeeLoginError{
		// 	Code:   400,
		// 	ErrMsg: "用户不存在",
		// })
	} else {
		err := bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(password))
		fmt.Print(err, "err")
		if err != nil {
			c.JSON(http.StatusBadRequest, vo.EmployeeLoginError{
				Code:   400,
				ErrMsg: "密码不正确",
			})
			return
		}
		// 当前用户是否存在，密码是否正确
		token, errToken := utils.GenerateAdminToken(config.Cfg.JWT.AdminSecret, config.Cfg.JWT.AdminTTL, employee.ID)
		if errToken != nil {
			c.JSON(http.StatusBadRequest, vo.EmployeeLoginError{
				Code:   400,
				ErrMsg: "生成Token失败",
			})
			return
		}
		c.JSON(http.StatusOK, vo.EmployeeLoginResponseVO{
			Code: 1,
			Data: vo.EmployeeLoginVO{
				ID:       employee.ID,
				Name:     employee.Name,
				Username: employee.Name,
				Token:    token,
			},
		})
	}

}
func Logout(c *gin.Context) { utils.Error(c, "not implemented") }
func AddEmployee(c *gin.Context) {
	var request dto.EmployeeAddDTO
	if error := c.ShouldBindJSON(&request); error != nil {
		c.JSON(http.StatusOK, vo.EmployeeLoginError{
			Code:   400,
			ErrMsg: "参数错误",
		})
		return
	}
	userName := request.Username
	employeeDao := dao.NewEmployeeDAO()
	employee, _ := employeeDao.FindByUserName(userName)
	if employee != nil {
		c.JSON(http.StatusOK, vo.EmployeeLoginError{
			Code:   400,
			ErrMsg: "管理员已存在，请更改名字重试",
		})
		return
	}
	employeeReq := dto.EmployeeAddDTO{
		Name:     request.Name,
		Username: request.Username,
		Sex:      request.Sex,
		Phone:    request.Phone,
	}
	result, error := employeeDao.CreateUserByModel(&employeeReq)
	if error != nil {
		c.AbortWithStatusJSON(http.StatusOK, vo.EmployeeLoginError{
			Code:   400,
			ErrMsg: "创建失败，请重试！",
		})
		return
	}

	c.JSON(http.StatusOK, vo.EmployeeResponseVO{
		Code: 1,
		Data: vo.EmployeeUserVO{
			ID:       result.ID,
			UserName: result.Username,
		},
	})
}
func EmployeePage(c *gin.Context) {
	var request dto.EmployeePageQueryDTO

	if error := c.ShouldBindQuery(&request); error != nil {
		c.JSON(http.StatusOK, vo.EmployeeLoginError{
			Code:   400,
			ErrMsg: "参数错误，请检查后重试",
		})
		return
	}
	employeeRequest := dto.EmployeePageQueryDTO{
		Name:     request.Name,
		Page:     request.Page,
		PageSize: request.PageSize,
	}
	employeeDao := dao.NewEmployeeDAO()
	employeeList, total, err := employeeDao.FindAllUser(&employeeRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, vo.EmployeeLoginError{
			Code:   400,
			ErrMsg: "查询员工失败",
		})
		return
	}
	records := make([]vo.EmployRecord, 0, len(employeeList))
	for _, list := range employeeList {
		records = append(records, vo.EmployRecord{
			Name:       list.Name,
			Username:   list.Username,
			ID:         list.ID,
			Phone:      list.Phone,
			Sex:        list.Sex,
			IDNumber:   list.IDNumber,
			Status:     list.Status,
			CreateTime: list.CreateTime,
			UpdateTime: list.UpdateTime,
			UpdateUser: list.UpdateUser,
			CreateUser: list.CreateUser,
		})
	}
	c.JSON(http.StatusOK, vo.EmployeeListResponseVO{
		Code: 1,
		Data: vo.EmployeeRecordsVO{
			Records: records,
			Total:   total,
		},
	})
}
func GetEmployee(c *gin.Context) {
	var request dto.EmployeeQueryById
	if error := c.ShouldBindUri(&request); error != nil {
		fmt.Print(error)
		c.AbortWithStatusJSON(http.StatusOK, vo.EmployeeLoginError{
			Code:   400,
			ErrMsg: "参数错误，请检查～",
		})
		return
	}

	employeeDao := dao.NewEmployeeDAO()
	employee, error := employeeDao.FindByUserId(request.ID)
	if error != nil {
		c.AbortWithStatusJSON(http.StatusOK, vo.EmployeeLoginError{
			Code:   400,
			ErrMsg: "查询失败，请重试～",
		})
		return
	}

	c.JSON(http.StatusOK, vo.EmployeeEditResponseVO{
		Code: 1,
		Data: vo.EmployRecord{
			Name:       employee.Name,
			Username:   employee.Username,
			ID:         employee.ID,
			Phone:      employee.Phone,
			Sex:        employee.Sex,
			IDNumber:   employee.IDNumber,
			Status:     employee.Status,
			CreateTime: employee.CreateTime,
			UpdateTime: employee.UpdateTime,
			UpdateUser: employee.UpdateUser,
			CreateUser: employee.CreateUser,
		},
	})
}
func UpdateEmployee(c *gin.Context) {
	var request dto.EmployeeUpdateDTO
	if error := c.ShouldBindJSON(&request); error != nil {
		c.AbortWithStatusJSON(http.StatusOK, vo.EmployeeLoginError{
			Code:   400,
			ErrMsg: "参数错误，请检查～",
		})
		return
	}

	employeeEdit := dto.EmployeeUpdateDTO{
		ID:       request.ID,
		IDNumber: request.IDNumber,
		Sex:      request.Sex,
		Phone:    request.Phone,
		Name:     request.Name,
		Username: request.Username,
	}
	employeeDao := dao.NewEmployeeDAO()
	error := employeeDao.EditUser(&employeeEdit)
	if error != nil {
		c.AbortWithStatusJSON(http.StatusOK, vo.EmployeeLoginError{
			Code:   400,
			ErrMsg: "编辑失败，请重试",
		})
		return
	}

	c.JSON(http.StatusOK, vo.EmployeeResponseVO{
		Code: 1,
		Data: vo.EmployeeUserVO{
			UserName: employeeEdit.Username,
			ID:       employeeEdit.ID,
		},
	})
}
func SetEmployeeStatus(c *gin.Context) {
	statusStr := c.Param("status")
	idStr := c.Query("id")

	status, err := strconv.Atoi(statusStr)
	if err != nil {
		c.JSON(http.StatusOK, vo.EmployeeLoginError{
			Code:   400,
			ErrMsg: "参数错误，请检查后重试",
		})
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, vo.EmployeeLoginError{
			Code:   400,
			ErrMsg: "参数错误，请检查后重试",
		})
		return
	}
	dao := dao.NewEmployeeDAO()
	result, err := dao.UpdateUserStatusById(id, status)
	if err != nil {
		fmt.Print(err, "error")
		c.JSON(http.StatusOK, vo.EmployeeLoginError{
			Code:   400,
			ErrMsg: "更新失败",
		})
		return
	}
	c.JSON(http.StatusOK, vo.EmployeeEditResponseVO{
		Code: 200,
		Data: vo.EmployRecord{
			ID: result,
		},
	})
}
func EditPassword(c *gin.Context) {
	var req dto.EmployeeEditPasswordDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, vo.EmployeeLoginError{
			Code:   400,
			ErrMsg: "参数错误，请检查～",
		})
		return
	}

	emp_id, exist := c.Get("emp_id")
	if !exist {
		c.AbortWithStatusJSON(http.StatusOK, vo.EmployeeLoginError{
			Code:   400,
			ErrMsg: "token失效，请重试～",
		})
		return
	}
	empDao := dao.NewEmployeeDAO()
	result, err := empDao.FindByUserId(emp_id.(uint64))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, vo.EmployeeLoginError{
			Code:   400,
			ErrMsg: "当前用户未查询到～",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(req.OldPassword)); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, vo.EmployeeLoginError{
			Code:   400,
			ErrMsg: "原密码错误",
		})
		return
	}

	err = empDao.UpdatePwdByEmpId(emp_id.(uint64), req.NewPassword)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, vo.EmployeeLoginError{
			Code:   400,
			ErrMsg: "修改失败",
		})
		return
	}

	c.JSON(http.StatusOK, vo.EmployeeResponseVO{
		Code: 200,
		Data: vo.EmployeeUserVO{
			ID: emp_id.(uint64),
		},
	})

}

// Category
func AddCategory(c *gin.Context) {
	var req dto.CategoryAddDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Print(err, "err")
		c.AbortWithStatusJSON(http.StatusOK, vo.CategoryResError{
			Code:   400,
			ErrMsg: "参数错误",
		})
		return
	}

	resultDao := dao.NewCategoryDAO()
	_, err := resultDao.FindCategoryByName(req.Name)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusOK, vo.CategoryResError{
			Code:   400,
			ErrMsg: "当前种类已存在",
		})
		return
	}

	var reqEntity = dto.CategoryAddDTO{
		Name: req.Name,
		Sort: req.Sort,
		Type: req.Type,
	}
	categoryID, err := resultDao.InsertCategory(&reqEntity)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, vo.CategoryResError{
			Code:   400,
			ErrMsg: "添加失败请重试～",
		})
		return
	}
	c.JSON(http.StatusOK, vo.CategoryResSuccess{
		Code: 200,
		Data: vo.CatorySuccessData{
			ID: categoryID,
		},
	})

}
func CategoryPage(c *gin.Context) {
	var req dto.CategoryPageQueryDTO
	if err := c.ShouldBindQuery(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, vo.CategoryResError{
			Code:   400,
			ErrMsg: "参数错误",
		})
		return
	}
	categoryDao := dao.NewCategoryDAO()
	result, total, err := categoryDao.FindAllCategory(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, vo.CategoryResError{
			Code:   400,
			ErrMsg: "查询失败",
		})
		return
	}

	c.JSON(http.StatusOK, vo.CategoryResListSuccess{
		Code: 200,
		Data: vo.CategoryListSuccessData{
			Records: result,
			Total:   total,
		},
	})
}
func DeleteCategory(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, vo.EmployeeLoginError{
			Code:   400,
			ErrMsg: "参数错误",
		})
		return
	}
	cateDao := dao.NewCategoryDAO()
	err = cateDao.DeleteCategoryById(id)
	if err != nil {
		c.JSON(http.StatusOK, vo.EmployeeLoginError{
			Code:   400,
			ErrMsg: "删除失败",
		})
	}

	c.JSON(http.StatusOK, vo.CategoryResSuccess{
		Code: 1,
	})
}
func UpdateCategory(c *gin.Context) {
	var req dto.CategoryUpdateDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Print(err, "err")
		c.AbortWithStatusJSON(http.StatusOK, vo.CategoryResError{
			Code:   400,
			ErrMsg: "参数错误",
		})
		return
	}

	resultDao := dao.NewCategoryDAO()
	sort := req.Sort
	err := resultDao.EditCategoryById(req.ID, req.Name, sort)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, vo.CategoryResError{
			Code:   400,
			ErrMsg: "更新失败请重试～",
		})
		return
	}
	c.JSON(http.StatusOK, vo.CategoryResSuccess{
		Code: 200,
		Data: vo.CatorySuccessData{
			ID: req.ID,
		},
	})
}
func SetCategoryStatus(c *gin.Context) {
	idStr := c.Query("id")
	statusStr := c.Param("status")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, vo.EmployeeLoginError{
			Code:   400,
			ErrMsg: "参数错误",
		})
		return
	}
	status, err := strconv.ParseUint(statusStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, vo.EmployeeLoginError{
			Code:   400,
			ErrMsg: "参数错误",
		})
		return
	}
	cateDao := dao.NewCategoryDAO()
	err = cateDao.UpdateCategoryById(id, status)
	if err != nil {
		c.JSON(http.StatusOK, vo.EmployeeLoginError{
			Code:   400,
			ErrMsg: "更新失败",
		})
		return
	}

	c.JSON(http.StatusOK, vo.CategoryResSuccess{
		Code: 1,
	})
}
func CategoryList(c *gin.Context) {
	var req dto.CategoryListQueryDTO
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusOK, vo.EmployeeLoginError{
			Code:   400,
			ErrMsg: "参数错误",
		})
		return
	}

	categoryDao := dao.NewCategoryDAO()
	categoryList, err := categoryDao.FindAllCategoryByType(req.Type)
	if err != nil {
		c.JSON(http.StatusOK, vo.EmployeeLoginError{
			Code:   400,
			ErrMsg: "查找失败",
		})

		return
	}

	c.JSON(http.StatusOK, vo.CategoryListResData{
		Code: 200,
		Data: categoryList,
	})
}

// Dish
func AddDish(c *gin.Context) {
	var req dto.DishAddDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, "参数错误～")
		return
	}

	dishDB := dao.NewDishDAO()
	result, err := dishDB.AddDish(&req)
	if err != nil {
		utils.Error(c, "添加失败")
		return
	}

	utils.Success(c, result)

}
func DishPage(c *gin.Context) {
	var req dto.DishPageQueryDTO
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.Error(c, "参数错误～")
		return
	}
	if statusStr, ok := c.GetQuery("status"); !ok || statusStr == "" {
		req.Status = nil
	}
	dishDB := dao.NewDishDAO()
	result, total, err := dishDB.FindAllDish(&req)
	if err != nil {
		utils.Error(c, "查询失败～")
		return
	}
	utils.SuccPage(c, result, total)
}
func DeleteDish(c *gin.Context) {
	idsStr := c.Query("ids")
	if idsStr == "" {
		utils.Error(c, "传参错误～")
		return
	}
	ids := strings.Split(idsStr, ",")
	uintIds := make([]uint64, len(ids))
	for index, id := range ids {
		result, _ := strconv.ParseUint(id, 10, 64)
		uintIds[index] = result
	}
	dishDao := dao.NewDishDAO()
	err := dishDao.DeleteDishByIds(uintIds)
	if err != nil {
		utils.Error(c, "删除失败～")
		return
	}

	utils.Success(c, "删除成功")
}
func GetDish(c *gin.Context) {
	var idStr = c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.Error(c, "参数错误")
	}

	dishDAO := dao.NewDishDAO()
	result, err := dishDAO.FindDishById(id)
	if err != nil {
		utils.Error(c, "查询失败～")
		return
	}

	utils.Success(c, result)

}
func UpdateDish(c *gin.Context) {
	var req dto.DishUpdateDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, "参数错误～")
		return
	}

	dishDB := dao.NewDishDAO()
	result, err := dishDB.UpdateDish(&req)
	if err != nil {
		utils.Error(c, "更新失败")
		return
	}

	utils.Success(c, result)
}
func SetDishStatus(c *gin.Context) {
	var req dto.DishStatusUpdateDTO
	if err := c.ShouldBind(&req); err != nil {
		utils.Error(c, "参数错误")
		return
	}

	dishDao := dao.NewDishDAO()
	err := dishDao.UpdateStatusById(req.Status, req.ID)

	if err != nil {
		utils.Error(c, "更新失败")
		return
	}

	utils.Success(c, &req)

}
func DishList(c *gin.Context) {
	var categoryIdStr = c.Query("categoryId")
	categoryId, err := strconv.ParseUint(categoryIdStr, 10, 64)
	if err != nil {
		utils.Error(c, "参数错误")
		return
	}

	var dishDao = dao.NewDishDAO()
	result, err := dishDao.FindDishBycategoryId(categoryId)
	if err != nil {
		utils.Error(c, "查询失败")
		return
	}

	utils.Success(c, result)
}

// Setmeal
func AddSetmeal(c *gin.Context) {
	var req dto.SetmealAddDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, "参数错误～")
		fmt.Println(err, "err")
		return
	}
	fmt.Println(req.Price)

	dishDao := dao.NewSetmealDAO()
	result, err := dishDao.AddSetmeal(&req)
	if err != nil {
		utils.Error(c, "增加失败～")
		return
	}

	utils.Success(c, result)
}
func SetmealPage(c *gin.Context) {
	var req dto.SetmealPageQueryDTO
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.Error(c, "参数错误～")
		return
	}

	if status, ok := c.GetQuery("status"); !ok || status == "" {
		req.Status = nil
	}
	setmealDao := dao.NewSetmealDAO()
	setmealList, total, err := setmealDao.FindAllSetmeal(&req)

	if err != nil {
		utils.Error(c, "查询失败")
		return
	}

	utils.SuccPage(c, setmealList, total)
}
func DeleteSetmeal(c *gin.Context) {
	idsStr := c.Query("ids")
	if idsStr == "" {
		utils.Error(c, "传参错误～")
		return
	}
	ids := strings.Split(idsStr, ",")
	uintIds := make([]uint64, len(ids))
	for index, id := range ids {
		result, _ := strconv.ParseUint(id, 10, 64)
		uintIds[index] = result
	}
	setmealDao := dao.NewSetmealDAO()
	err := setmealDao.DeleteSetmealByIds(uintIds)
	if err != nil {
		utils.Error(c, "删除失败～")
		return
	}

	utils.Success(c, "删除成功")
}
func GetSetmeal(c *gin.Context) {
	var idstr = c.Param("id")
	id, err := strconv.ParseUint(idstr, 10, 64)
	if err != nil {
		utils.Error(c, "参数错误～")
		return
	}

	setmealDao := dao.NewSetmealDAO()
	setmeal, sd, dishes, err := setmealDao.FindSetmealById(id)
	if err != nil {
		utils.Error(c, "查询失败")
		return
	}

	dishMap := make(map[uint64]entity.Dish)
	for _, d := range dishes {
		dishMap[d.ID] = d
	}

	vos := make([]vo.SetmealDishVO, len(sd))
	for i, item := range sd {
		d := dishMap[item.DishID]
		vos[i] = vo.SetmealDishVO{
			DishID: item.DishID,
			Name:   d.Name,
			Price:  d.Price,
			Copies: item.Copies,
		}
	}

	result := vo.SetmealVO{
		ID:            setmeal.ID,
		Name:          setmeal.Name,
		CategoryID:    setmeal.CategoryID,
		Price:         setmeal.Price,
		Image:         setmeal.Image,
		Description:   setmeal.Description,
		Status:        setmeal.Status,
		SetmealDishes: vos,
	}

	utils.Success(c, result)

}
func UpdateSetmeal(c *gin.Context) {
	var req dto.SetmealUpdateDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, "参数错误～")
		return
	}

	dishDB := dao.NewSetmealDAO()
	result, err := dishDB.UpdateSetmeal(&req)
	if err != nil {
		utils.Error(c, "更新失败")
		return
	}

	utils.Success(c, result)
}
func SetSetmealStatus(c *gin.Context) {
	var req dto.DishStatusUpdateDTO
	if err := c.ShouldBind(&req); err != nil {
		utils.Error(c, "参数错误")
		return
	}

	dishDao := dao.NewDishDAO()
	err := dishDao.UpdateStatusById(req.Status, req.ID)

	if err != nil {
		utils.Error(c, "更新失败")
		return
	}

	utils.Success(c, &req)
}
func SetmealList(c *gin.Context)      { utils.Error(c, "not implemented") }
func GetSetmealDishes(c *gin.Context) { utils.Error(c, "not implemented") }

// Order
func OrderConditionSearch(c *gin.Context) {
	var req dto.OrderPageQueryDTO
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.Error(c, "参数错误")
		return
	}

	orderDao := dao.NewOrderDAO()
	list, total, err := orderDao.ConditionSearch(&req)
	if err != nil {
		utils.Error(c, "查询失败")
		return
	}

	utils.Success(c, gin.H{
		"records": list,
		"total":   total,
	})
}
func OrderStatistics(c *gin.Context) {
	orderDao := dao.NewOrderDAO()
	stats, err := orderDao.Statistics()
	if err != nil {
		utils.Error(c, "查询失败")
		return
	}

	// 状态映射为中文名
	type statVO struct {
		Name  string `json:"name"`
		Count int64  `json:"count"`
	}
	result := make([]statVO, len(stats))
	statusMap := map[int]string{
		utils.OrderStatusPendingPay:    "待付款",
		utils.OrderStatusPendingAccept: "待接单",
		utils.OrderStatusAccepted:      "已接单",
		utils.OrderStatusDelivering:    "配送中",
		utils.OrderStatusCompleted:     "已完成",
		utils.OrderStatusCancelled:     "已取消",
	}
	for i, s := range stats {
		result[i] = statVO{Name: statusMap[s.Status], Count: s.Count}
	}
	utils.Success(c, result)
}

func OrderDetails(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.Error(c, "参数错误")
		return
	}

	orderDao := dao.NewOrderDAO()
	order, err := orderDao.GetOrderByID(id)
	if err != nil {
		utils.Error(c, "订单不存在")
		return
	}
	utils.Success(c, order)
}
func ConfirmOrder(c *gin.Context) {
	var req dto.OrderConfirmDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, "参数错误")
		return
	}
	// userID := middleware.GetUserID(c)
	orderDao := dao.NewOrderDAO()
	if err := orderDao.ConfirmOrder(req.ID); err != nil {
		utils.Error(c, err.Error())
		return
	}
	msg, _ := json.Marshal(gin.H{
		"orderId":   req.ID,
		"eventName": "new_order",
	})
	utils.Send("order.dispatch", msg)
	utils.Success(c, nil)
}
func RejectOrder(c *gin.Context) {
	var req dto.OrderRejectionDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, "参数错误")
		return
	}
	if err := dao.NewOrderDAO().RejectOrder(req.ID, req.RejectionReason); err != nil {
		utils.Error(c, err.Error())
		return
	}
	utils.Success(c, nil)
}
func CancelOrder(c *gin.Context) {
	var req dto.OrderCancelDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, "参数错误")
		return
	}
	if err := dao.NewOrderDAO().CancelOrder(req.ID, req.CancelReason); err != nil {
		utils.Error(c, err.Error())
		return
	}
	utils.Success(c, nil)
}
func DeliveryOrder(c *gin.Context) {
	var req dto.OrderConfirmDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, "参数错误")
		return
	}
	if err := dao.NewOrderDAO().DeliveryOrder(req.ID); err != nil {
		utils.Error(c, err.Error())
		return
	}
	utils.Success(c, nil)
}
func CompleteOrder(c *gin.Context) {
	var req dto.OrderConfirmDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, "参数错误")
		return
	}
	if err := dao.NewOrderDAO().CompleteOrder(req.ID); err != nil {
		utils.Error(c, err.Error())
		return
	}
	utils.Success(c, nil)
}

// Shop
func GetShopStatus(c *gin.Context) {
	ctx := c.Request.Context()
	status, err := dao.Redis.Get(ctx, "SHOP_STATUS").Result()
	if err != nil {
		utils.Error(c, "获取营业状态失败")
		return
	}
	utils.Success(c, status)
}

func SetShopStatus(c *gin.Context) {
	statusStr := c.Param("status")
	status, err := strconv.Atoi(statusStr)
	if err != nil {
		utils.Error(c, "参数错误")
		return
	}
	ctx := c.Request.Context()
	if err := dao.Redis.Set(ctx, "SHOP_STATUS", status, 0).Err(); err != nil {
		utils.Error(c, "设置营业状态失败")
		return
	}
	utils.Success(c, nil)
}

// Common
func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	// files := form.File["files"]
	// if len(files) == 0 {
	// 	utils.Error(c, "no files uploaded")
	// 	return
	// }

	var paths []string
	name, err := utils.SaveFile(file)
	if err != nil {
		utils.Error(c, err.Error())
		return
	}
	paths = append(paths, name)
	utils.Success(c, vo.UpdateData{
		Name: name,
	}) // 返回文件名数组，前端拿到后写回 Image 字段
}

// Report
func TurnoverStatistics(c *gin.Context) { utils.Error(c, "not implemented") }
func UserStatistics(c *gin.Context)     { utils.Error(c, "not implemented") }
func OrdersStatistics(c *gin.Context)   { utils.Error(c, "not implemented") }
func Top10(c *gin.Context)              { utils.Error(c, "not implemented") }
func Export(c *gin.Context)             { utils.Error(c, "not implemented") }

// Workspace
func BusinessData(c *gin.Context) {
	data, err := dao.NewOrderDAO().BusinessData()
	if err != nil {
		utils.Error(c, "查询失败")
		return
	}
	utils.Success(c, data)
}
func OverviewOrders(c *gin.Context) {
	// 查询所有状态的Order
	var overviewDao = dao.NewOrderDAO()
	orderList, err := overviewDao.GetAllOrder()
	if err != nil {
		utils.Error(c, "查询失败")
		return
	}

	utils.Success(c, utils.FilterMap(orderList))

}
func OverviewDishes(c *gin.Context) {
	var overviewDao = dao.NewDishDAO()
	dishList, err := overviewDao.FindAllDishByALl()
	if err != nil {
		utils.Error(c, "查询失败")
		return
	}
	fmt.Printf("Dish count: " + strconv.Itoa(len(*dishList)))

	utils.Success(c, utils.FilterMapDish(*dishList))
}
func OverviewSetmeals(c *gin.Context) {
	var overviewDao = dao.NewSetmealDAO()
	setmealList, err := overviewDao.FindAllSetmealByALl()
	if err != nil {
		utils.Error(c, "查询失败")
		return
	}

	utils.Success(c, utils.FilterMapSetmeal(setmealList))
}
