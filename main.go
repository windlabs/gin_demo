package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"net/http"
	_ "windlabs.com/gin/docs" // 这里需要引入本地已生成文档
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}


// @title Swagger Example API12222
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8085
// @BasePath /api/v1
func main() {
	route := gin.Default()
	route.Any("/testing", startPage)
	route.Any("/check", connectCheck)
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := route.Group("/api/v1")
	{
		v1.GET("/test", Test)
	}
	route.Run(":8085")
}

type Response struct {
	Code    uint32      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseError struct {
	Code    uint32 `json:"code"`
	Message uint32 `json:"message"`
}


// @Summary 服务连接校验  接口简介
// @Description 服务初始连接测试  接口描述
// Success 200 {object} Response  成功后返回数据结构
// Failure 400 {object} ResponseError  失败后返回数据结构
// Failure 404 {object} ResponseError
// Failure 500 {object} ResponseError
// @Router /check [get] 路由地址及请求方法
func connectCheck(c *gin.Context) {
	res := Response{Code: 1001, Message: "OK", Data: "connect success !!!"}
	c.JSON(http.StatusOK, res)
}

// @summary 服务连接校验  接口简介
// @Description 服务初始连接测试  接口描述
// Success 200 {object} Response  成功后返回数据结构
// Failure 400 {object} ResponseError  失败后返回数据结构
// Failure 404 {object} ResponseError
// Failure 500 {object} ResponseError
// @Router /hello [get] 路由地址及请求方法
func startPage(c *gin.Context) {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "Success")
}
type Res struct {
	code int `json:"code"`
	message string `json:"message"`
}
// @Title 这是test标题
// @Summary 首页接口摘要信息
// @Description 接口描述信息
// @Tags 用户信息
// @Accept json
// @Security Bearer
// @Param name path string false "name value"
// @Success 200  "{"code":200, "message":"success"}"
// @Failure 400 {string} Res "{"code":400, "message":"success"}"
// @Failure 404 {string} Res "{"code":404, "message":"success"}"
// @Failure 500 {string} Res "{"code":500, "message":"success"}"
// @Router /test [get]
func Test(c *gin.Context) {
	c.JSON(200, Response{Code: 1001, Message: "OK", Data: "connect success !!!"})
}
