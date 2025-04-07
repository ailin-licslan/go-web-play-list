package routers

import (
	"github.com/gin-gonic/gin"
	"go-web-play/controller"
	"go-web-play/setting"
)

func SetupRouter() *gin.Engine {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default() //如果放开下面的 2个注释  不用依赖前端服务  已经打包为了静态文件放在后端服务一起了
	// 告诉gin框架模板文件引用的静态文件去哪里找
	//r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	//r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		// 修改某一个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
