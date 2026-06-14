package utils

import "github.com/gin-gonic/gin"

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ResultPage struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Total int64       `json:"total"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, Result{Code: CodeSuccess, Msg: "success", Data: data})
}

func Error(c *gin.Context, msg string) {
	c.JSON(200, Result{Code: CodeError, Msg: msg, Data: nil})
}

func SuccPage(c *gin.Context, data interface{}, total int64) {
	c.JSON(200, ResultPage{Code: CodeSuccess, Msg: "success", Data: data, Total: total})
}
