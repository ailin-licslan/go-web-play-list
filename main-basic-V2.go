package main

import (
	"github.com/gin-gonic/gin"
	"go-web-play/middleware"
	"log"
	"net/http"
	_ "net/http"
)

/*
*
GIN & ORM STUDY  参考 https://liwenzhou.com/posts/Go/gin/
including {goroutine middleware orm query upload  main-basic-V2.go}
*/
//func main() {   可以放开如果想学基础内容
func main2() {

	//basic study start ----
	r := gin.Default()

	//g.H  ====> map[string]interface{}
	//data := map[string]interface{}{"name": "this is a test", "age": 18, "hobby": "english"}

	r.GET("/LIN", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": "LIN",
		})
		//c.JSON(200, data)
	})

	//结构体写法
	type msg struct {
		Name  string
		Age   int
		Hobby string
	}
	r.GET("/LIN2", func(c *gin.Context) {
		c.JSON(200, msg{
			"LIN2",
			18,
			"English",
		})
	})

	//HTTP重定向
	r.GET("/TEST", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com/")
	})

	//路由重定向
	r.GET("/test", func(c *gin.Context) {
		// 指定重定向的URL
		c.Request.URL.Path = "/test2"
		r.HandleContext(c)
	})
	r.GET("/test2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"age": "18"})
	})

	//普通路由
	r.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": "LIN",
		})
	})
	r.GET("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": "LIN",
		})
	})
	r.POST("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": "LIN",
		})
	})
	r.Any("/111", func(c *gin.Context) { c.JSON(200, "123456") })

	r.LoadHTMLGlob("upload/*")
	r.Static("/upload", "upload")
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	//路由组
	userGroup := r.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) { c.JSON(200, "123456") })
		userGroup.GET("/login", func(c *gin.Context) { c.JSON(200, "123456") })
		userGroup.POST("/login", func(c *gin.Context) { c.JSON(200, "123456") })

	}
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) { c.JSON(200, "123456") })
		shopGroup.GET("/cart", func(c *gin.Context) { c.JSON(200, "123456") })
		shopGroup.POST("/checkout", func(c *gin.Context) { c.JSON(200, "123456") })
	}

	//路由组也是支持嵌套的
	goodsGroup := r.Group("/goods")
	{
		goodsGroup.GET("/index", func(c *gin.Context) { c.JSON(200, "123456") })
		goodsGroup.GET("/cart", func(c *gin.Context) { c.JSON(200, "123456") })
		goodsGroup.POST("/checkout", func(c *gin.Context) { c.JSON(200, "123456") })
		// 嵌套路由组
		xx := goodsGroup.Group("xx")
		xx.GET("/oo", func(c *gin.Context) { c.JSON(200, "123456") })
	}

	// 注册一个全局中间件 StatCost
	//r.Use(middleware.StatCost())

	// 给/test1111路由单独注册中间件（可注册多个）
	r.GET("/test1111", middleware.StatCost(), func(c *gin.Context) {
		name := c.MustGet("name").(string) // 从上下文取值
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!11111111111111111111111111",
		})
	})

	r.GET("/name", func(c *gin.Context) {
		name := c.MustGet("name").(string) // 从上下文取值
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})

	//为路由组注册中间件  2种形式

	shopGroupx := r.Group("/shopx", middleware.StatCost())
	{
		shopGroupx.GET("/index", func(c *gin.Context) { c.JSON(200, "123456123123212131111111111") })
	}

	shopGroupy := r.Group("/shopy")
	shopGroupy.Use(middleware.StatCost())
	{
		shopGroupy.GET("/index", func(c *gin.Context) {
			log.Println("cccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc")
			c.JSON(200, "123ccccccccccccccccc6123123212131111111111")
		})
	}

	//gin默认中间件
	//gin.Default()默认使用了Logger和Recovery中间件，其中：
	//Logger中间件将日志写入gin.DefaultWriter，即使配置了GIN_MODE=release。
	//Recovery中间件会recover任何panic。如果有panic的话，会写入500响应码。
	//如果不想使用上面两个默认的中间件，可以使用gin.New()新建一个没有任何默认中间件的路由。

	//basic study end ----

	_ = r.Run(":8080")

}
