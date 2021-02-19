package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var Router *gin.Engine

func GinInit()  {
	// 禁用控制台颜色
	//gin.DisableConsoleColor()

	//gin.New()返回一个*Engine 指针
	//而gin.Default()不但返回一个*Engine 指针，而且还进行了debugPrintWARNINGDefault()和engine.Use(Logger(), Recovery())其他的一些中间件操作
	Router = gin.Default()
	//Router = gin.New()
}

func SetupRouter(projectPath string) {

	//使用日志
	//Router.Use(gin.Logger())
	//使用Panic处理方案
	//Router.Use(gin.Recovery())

	Router.Use(InitErrorHandler)
	Router.Use(InitAccessLogMiddleware)

	// 未知调用方式
	Router.NoMethod(InitNoMethodJson)
	// 未知路由处理
	Router.NoRoute(InitNoRouteJson)

	// Ping
	Router.GET(projectPath + "/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ping": "pong",
		})
	})

	Router.POST(projectPath + "/pp", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ping": "post",
		})
	})

}
