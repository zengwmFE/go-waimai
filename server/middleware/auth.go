package middleware

import (
	"net/http"

	"sky-take-out-go/config"
	"sky-take-out-go/utils"

	"github.com/gin-gonic/gin"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenName := config.Cfg.JWT.AdminTokenName
		tokenStr := c.GetHeader(tokenName)
		if tokenStr == "" {
			c.AbortWithStatusJSON(http.StatusOK, utils.Result{Code: utils.CodeError, Msg: "NOT_LOGIN", Data: nil})
			return
		}
		claims, err := utils.ParseAdminToken(config.Cfg.JWT.AdminSecret, tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, utils.Result{Code: utils.CodeError, Msg: "TOKEN_INVALID", Data: nil})
			return
		}
		c.Set("emp_id", claims.EmpID)
		c.Next()
	}
}

func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenName := config.Cfg.JWT.UserTokenName
		tokenStr := c.GetHeader(tokenName)
		if tokenStr == "" {
			c.AbortWithStatusJSON(http.StatusOK, utils.Result{Code: utils.CodeError, Msg: "NOT_LOGIN", Data: nil})
			return
		}
		claims, err := utils.ParseUserToken(config.Cfg.JWT.UserSecret, tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, utils.Result{Code: utils.CodeError, Msg: "TOKEN_INVALID", Data: nil})
			return
		}
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}

func GetEmpID(c *gin.Context) uint64 {
	id, _ := c.Get("emp_id")
	return id.(uint64)
}

func GetUserID(c *gin.Context) uint64 {
	id, _ := c.Get("user_id")
	return id.(uint64)
}
