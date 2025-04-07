package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//推荐使用社区的https://github.com/gin-contrib/cors 库，一行代码解决前后端分离架构下的跨域问题。
//注意： 该中间件需要注册在业务处理函数前面。
//这个库支持各种常用的配置项，具体使用方法如下。

func main() {
	router := gin.Default()
	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},                         // 允许跨域发来请求的网站
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 允许的请求方法
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool { // 自定义过滤源站的方法
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	router.Run()
}

//为全局路由注册

//也可以简单的像下面的示例代码那样使用默认配置，允许所有的跨域请求
//func main() {
//	router := gin.Default()
//	// same as
//	// config := cors.DefaultConfig()
//	// config.AllowAllOrigins = true
//	// router.Use(cors.New(config))
//	router.Use(cors.Default())
//	router.Run()
//}
