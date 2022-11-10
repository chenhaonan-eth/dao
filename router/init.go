package router

import (
	"net/http"

	"github.com/chenhaonan-eth/guide-sisyphean/core"
	"github.com/chenhaonan-eth/guide-sisyphean/router/api"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var get = map[string]gin.HandlerFunc{
	"/getEquityBbondYieldSpreads": api.GetdEquityBbondYieldSpreads,
	"/ping":                       ping,
}

var post = map[string]gin.HandlerFunc{
	// "/topics/:topic_name/:partition": api.GetTopicData,
}

func ping(c *gin.Context) {
	core.G_LOG.Info("a sample app log")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				core.G_LOG.Error("Panic info is: ", zap.Any("err", err))
			}
		}()
		c.Next()
	}
}
func WebInit(router *gin.Engine) {
	//react npm run build
	// router.LoadHTMLGlob("./dist/*.html") // npm打包成dist的路径
	// router.Static("/favicon.ico", "./dist/favicon.ico")
	// router.Static("/static", "./dist/static")   // dist里面的静态资源
	// router.StaticFile("/", "./dist/index.html") // 前端网页入口页面

	router.Use(Cors()) //开启中间件 允许使用跨域请求

	for path, handler := range get {
		router.GET(path, handler)
	}

	for path, handler := range post {
		router.POST(path, handler)
	}
}
