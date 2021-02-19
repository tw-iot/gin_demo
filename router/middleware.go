package router

import (
	"encoding/json"
	. "gin_demo/log"
	. "gin_demo/threadlocal"
	"github.com/gin-gonic/gin"
	. "github.com/jtolds/gls"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// ErrorHandler is a middleware to handle errors encountered during requests
func InitErrorHandler(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": c.Errors,
		})
	}
}

//未知路由处理 返回json
func InitNoRouteJson(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"code": http.StatusNotFound,
		"msg":  "path not found",
	})
}

//未知调用方式 返回json
func InitNoMethodJson(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, gin.H{
		"code": http.StatusMethodNotAllowed,
		"msg":  "method not allowed",
	})
}

//打印请求和响应日志
func InitAccessLogMiddleware(c *gin.Context) {
	//request id
	requestId := c.Request.Header.Get("X-RequestId")
	if requestId == "" {
		requestId = strconv.FormatInt(time.Now().UnixNano(), 10)
	}
	//response requestId
	c.Writer.Header().Set("X-RequestId", requestId)

	// 开始时间
	startTime := time.Now()

	//处理请求 do chian
	Mgr.SetValues(Values{Rid: requestId}, func() {
		c.Next()
	})

	// 结束时间
	endTime := time.Now()
	// 执行时间
	latencyTime := endTime.Sub(startTime)
	// 请求方式
	reqMethod := c.Request.Method
	// 请求路由
	reqUri := c.Request.RequestURI
	// 状态码
	statusCode := c.Writer.Status()
	// 请求IP
	clientIP := c.ClientIP()
	//请求参数
	body, _ := ioutil.ReadAll(c.Request.Body)
	//返回参数
	responseMap := c.Keys
	responseJson, _ := json.Marshal(responseMap)

	//日志格式
	//LogAccess.Infof("| %3d | %13v | %15s | %s | %s | %s | %s | %s |",
	//	statusCode,
	//	latencyTime,
	//	clientIP,
	//	reqMethod,
	//	reqUri,
	//	requestId,
	//	string(body),
	//	string(responseJson),
	//)

	// 日志格式
	LogAccess.WithFields(logrus.Fields{
		"status_code":  statusCode,
		"latency_time": latencyTime,
		"client_ip":    clientIP,
		"req_method":   reqMethod,
		"req_uri":      reqUri,
		"req_Id":       requestId,
		"req_body":     string(body),
		"res_body":     string(responseJson),
	}).Info()

}
