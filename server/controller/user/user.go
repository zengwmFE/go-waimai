package user

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sky-take-out-go/config"
	"sky-take-out-go/dao"
	"sky-take-out-go/middleware"
	"sky-take-out-go/model/dto"
	"sky-take-out-go/model/vo"
	"sky-take-out-go/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WxLoginResponse struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
	UnionId    string `json:"unionid"`
}

// User
func Login(c *gin.Context) {
	// https://api.weixin.qq.com/sns/jscode2session
	var request dto.UserLoginDTO

	if error := c.ShouldBindJSON(&request); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}

	code := request.Code
	u, _ := url.Parse("https://api.weixin.qq.com/sns/jscode2session")
	u.RawQuery = url.Values{
		"js_code":    {code},
		"appid":      {"wx8da70836f8424e6b"},
		"secret":     {"20ccacd613a901c5fd7997afd55c0af5"},
		"grant_type": {"authorization_code"},
	}.Encode()
	response, error := http.Get(u.String())
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Read response body error:", err)
		return
	}
	var wxResp WxLoginResponse
	err = json.Unmarshal(body, &wxResp)

	if err != nil {
		fmt.Println("JSON unmarshal error:", err)
		return
	}
	if wxResp.ErrCode != 0 {
		fmt.Println("WeChat API error:", wxResp.ErrMsg)
		return
	}
	// 判断用户是否存在
	userDAO := dao.NewUserDAO()
	user, _ := userDAO.FindByUserId(wxResp.Openid)
	if err != nil {
		utils.Error(c, "创建token失败")
		return
	}
	if user == nil {
		// 不存在，那么就获取用户微信信息 -> DAO新建用户
		newUser, err := userDAO.CreateUserId(wxResp.Openid)
		if err != nil {
			utils.Error(c, "创建失败")
			return
		} else {
			token, err := utils.GenerateUserToken(config.Cfg.JWT.UserSecret, config.Cfg.JWT.UserTTL, newUser.ID)
			if err != nil {
				utils.Error(c, "创建token失败")
				return
			}
			utils.Success(c, vo.UserLoginVO{
				ID:    newUser.ID,
				Name:  newUser.Name,
				Token: token,
			})
		}
	} else {
		token, err := utils.GenerateUserToken(config.Cfg.JWT.UserSecret, config.Cfg.JWT.UserTTL, user.ID)
		if err != nil {
			utils.Error(c, "创建token失败")
			return
		}
		utils.Success(c, vo.UserLoginVO{
			ID:    user.ID,
			Name:  user.Name,
			Token: token,
		})
	}
}
func Logout(c *gin.Context) { utils.Error(c, "not implemented") }

// Category
func CategoryList(c *gin.Context) {
	typeStr := c.Query("type")

	typeID, err := strconv.Atoi(typeStr)
	if err != nil {
		utils.Error(c, "参数错误")
		return
	}

	categoryDao := dao.NewCategoryDAO()
	result, err := categoryDao.FindAllCategoryByType(typeID)
	if err != nil {
		utils.Error(c, "查询失败")
		return
	}
	utils.Success(c, result)
}

// Dish
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
func SetmealList(c *gin.Context)      { utils.Error(c, "not implemented") }
func GetSetmealDishes(c *gin.Context) { utils.Error(c, "not implemented") }

// Shop
func GetShopStatus(c *gin.Context) {
	ctx := c.Request.Context()
	statusStr, err := dao.Redis.Get(ctx, "SHOP_STATUS").Result()
	if err != nil {
		utils.Error(c, "获取营业状态失败")
		return
	}
	status, err := strconv.Atoi(statusStr)
	utils.Success(c, status)
}

// ShoppingCart
func AddToCart(c *gin.Context) {
	var req dto.ShoppingCartAddDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, "参数错误")
		return
	}
	if req.DishID == nil && req.SetmealID == nil {
		utils.Error(c, "请选择菜品或套餐")
		return
	}

	userID := middleware.GetUserID(c)
	if err := dao.NewShoppingCartDAO().AddToCart(userID, req); err != nil {
		utils.Error(c, "添加购物车失败")
		return
	}
	utils.Success(c, nil)
}

func CartList(c *gin.Context) {
	userID := middleware.GetUserID(c)
	list, err := dao.NewShoppingCartDAO().GetCartList(userID)
	if err != nil {
		utils.Error(c, "查询购物车失败")
		return
	}
	utils.Success(c, list)
}

func CleanCart(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if err := dao.NewShoppingCartDAO().CleanCart(userID); err != nil {
		utils.Error(c, "清空购物车失败")
		return
	}
	utils.Success(c, nil)
}

func SubFromCart(c *gin.Context) {
	var req dto.ShoppingCartSubDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, "参数错误")
		return
	}

	userID := middleware.GetUserID(c)
	if err := dao.NewShoppingCartDAO().SubFromCart(userID, req); err != nil {
		utils.Error(c, "操作失败")
		return
	}
	utils.Success(c, nil)
}

// Order
func SubmitOrder(c *gin.Context) {
	var req dto.OrderSubmitDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, "参数错误～")
		return
	}

	// // 事务：创建订单 + 清购物车
	orderDao := dao.NewOrderDAO()
	userID := middleware.GetUserID(c)

	order, err := orderDao.SubmitOrder(userID, req)
	if err != nil {
		return
	}
	utils.Success(c, order)
}
func Payment(c *gin.Context) {
	var req dto.OrderPaymentDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, "参数错误")
		return
	}

	userID := middleware.GetUserID(c)
	if err := dao.NewOrderDAO().PayOrder(userID, req.OrderNumber, req.PayMethod); err != nil {
		utils.Error(c, err.Error())
		return
	}
	utils.Success(c, nil)
}
func HistoryOrders(c *gin.Context)   { utils.Error(c, "not implemented") }
func OrderDetail(c *gin.Context)     { utils.Error(c, "not implemented") }
func CancelOrder(c *gin.Context)     { utils.Error(c, "not implemented") }
func RepetitionOrder(c *gin.Context) { utils.Error(c, "not implemented") }
func Reminder(c *gin.Context)        { utils.Error(c, "not implemented") }

// AddressBook
func AddAddress(c *gin.Context) {
	var req dto.AddressBookAddDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, "参数错误")
		return
	}

	userID := middleware.GetUserID(c)
	dao := dao.NewAddressBookDAO()
	id, err := dao.AddAddress(userID, req)
	if err != nil {
		utils.Error(c, "新增地址失败")
		return
	}
	utils.Success(c, id)
}

func UpdateAddress(c *gin.Context) {
	var req dto.AddressBookUpdateDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, "参数错误")
		return
	}

	userID := middleware.GetUserID(c)
	err := dao.NewAddressBookDAO().UpdateAddress(userID, req)
	if err != nil {
		utils.Error(c, "更新地址失败")
		return
	}
	utils.Success(c, nil)
}

func AddressList(c *gin.Context) {
	userID := middleware.GetUserID(c)
	result, err := dao.NewAddressBookDAO().GetAddressList(userID)
	if err != nil {
		utils.Error(c, "查询失败")
		return
	}
	utils.Success(c, result)
}

func GetAddress(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.Error(c, "参数错误")
		return
	}

	userID := middleware.GetUserID(c)
	addr, err := dao.NewAddressBookDAO().GetAddressByID(id, userID)
	if err != nil {
		utils.Error(c, "地址不存在")
		return
	}
	utils.Success(c, addr)
}

func DeleteAddress(c *gin.Context) {
	var req dto.AddressBookDeleteDTO
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.Error(c, "参数错误")
		return
	}

	userID := middleware.GetUserID(c)
	if err := dao.NewAddressBookDAO().DeleteAddress(req.ID, userID); err != nil {
		utils.Error(c, "删除失败")
		return
	}
	utils.Success(c, nil)
}

func SetDefaultAddress(c *gin.Context) {
	var req dto.AddressBookDeleteDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, "参数错误")
		return
	}

	userID := middleware.GetUserID(c)
	if err := dao.NewAddressBookDAO().SetDefaultAddress(req.ID, userID); err != nil {
		utils.Error(c, "设置默认地址失败")
		return
	}
	utils.Success(c, nil)
}

func GetDefaultAddress(c *gin.Context) {
	userID := middleware.GetUserID(c)
	addr, err := dao.NewAddressBookDAO().GetDefaultAddress(userID)
	if err != nil {
		utils.Error(c, "暂无默认地址")
		return
	}
	utils.Success(c, addr)
}
