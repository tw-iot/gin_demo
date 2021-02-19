package main

import (
	"flag"
	"fmt"
	. "gin_demo/config"
	_ "gin_demo/docs"
	. "gin_demo/log"
	"gin_demo/router"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"runtime"
	"time"
)

var version = flag.Bool("version", true, "是否打印版本,默认打印")
var swagger = flag.Bool("swagger", true, "是否启动swagger接口文档,默认不启动")
var configFile = flag.String("configFile", "config/config.yml", "配置文件路径")
var projectPath = flag.String("projectPath", "/gin_demo", "项目访问路径前缀")

func init(){
	flag.Parse()

	ConfigRead(*configFile)

	LogInit()
}

//@title gin示例 API
//@version 0.0.1
//@description  相关接口文档
//@host 127.0.0.1:8080
//@BasePath
func main() {
	if *version {
		showVersion := fmt.Sprintf("%s %s@%s", "gin_demo", "1.0.0", time.Now().Format("2006-01-02 15:04:05"))
		fmt.Println(showVersion)
		fmt.Println("go version: " + runtime.Version())
	}

	Log.Info("start server...")

	gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	router.GinInit()

	//gin工程实例 *gin.Engine
	r := router.Router

	//路由初始化
	router.SetupRouter(*projectPath)

	if *swagger {
		//启动访问swagger文档
		r.GET(*projectPath + "/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	Log.Info("listen on :%s", Cfg.ListenPort)
	//监听端口
	r.Run(":" + Cfg.ListenPort)

}
