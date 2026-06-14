package rider

import (
	"net/http"
	"sky-take-out-go/config"
	"sky-take-out-go/dao"
	"sky-take-out-go/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	Phone string `json:"phone" binding:"required"`
}

func Login(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, "请输入手机号")
		return
	}

	rider, err := dao.NewRiderDAO().FindByPhone(req.Phone)
	if err != nil {
		utils.Error(c, "骑手不存在")
		return
	}

	token, err := utils.GenerateAdminToken(config.Cfg.JWT.AdminSecret, config.Cfg.JWT.AdminTTL, rider.ID)
	if err != nil {
		utils.Error(c, "生成Token失败")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": gin.H{
			"id":    rider.ID,
			"name":  rider.Name,
			"phone": rider.Phone,
			"token": token,
		},
	})
}

func PendingOrders(c *gin.Context) {
	orders, err := dao.NewRiderDAO().GetPendingOrders()
	if err != nil {
		utils.Error(c, "查询失败")
		return
	}
	utils.Success(c, orders)
}

func GrabOrder(c *gin.Context) {
	orderID, err := strconv.ParseUint(c.Query("id"), 10, 64)
	if err != nil {
		utils.Error(c, "参数错误")
		return
	}

	riderID := c.GetUint64("emp_id") // 复用 admin JWT，emp_id = rider_id
	if err := dao.NewRiderDAO().GrabOrder(orderID, riderID); err != nil {
		utils.Error(c, err.Error())
		return
	}
	utils.Success(c, nil)
}

func MyOrders(c *gin.Context) {
	riderID := c.GetUint64("emp_id")
	orders, err := dao.NewRiderDAO().GetMyOrders(riderID)
	if err != nil {
		utils.Error(c, "查询失败")
		return
	}
	utils.Success(c, orders)
}

func CompleteOrder(c *gin.Context) {
	orderID, err := strconv.ParseUint(c.Query("id"), 10, 64)
	if err != nil {
		utils.Error(c, "参数错误")
		return
	}

	riderID := c.GetUint64("emp_id")
	if err := dao.NewRiderDAO().CompleteOrder(orderID, riderID); err != nil {
		utils.Error(c, err.Error())
		return
	}
	utils.Success(c, nil)
}
