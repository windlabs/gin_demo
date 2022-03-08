package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct{
	Code uint32 `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type ResponseError struct{
	Code uint32 `json:"code"`
	Message uint32 `json:"message"`
}

// @summary 服务连接校验 --> 接口简介
// @Description 服务初始连接测试 --> 接口描述
// @Accept json   --> 接收类型
// @Produce json  --> 返回类型
// Success 200 {object} Response --> 成功后返回数据结构
// Failure 400 {object} ResponseError --> 失败后返回数据结构
// Failure 404 {object} ResponseError
// Failure 500 {object} ResponseError
// @Router /check [get] --> 路由地址及请求方法

func ConnectCheck(c *gin.Context){
	res := Response{ Code: 1001, Message: "OK", Data: "connect success !!!"}
	c.JSON(http.StatusOK, res)
}
