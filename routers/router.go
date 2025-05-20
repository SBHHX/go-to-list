package routers

import (
	"github.com/gin-gonic/gin"
	"go_to_list/controllers"
)

// 配置路由
func SetupRouter() *gin.Engine {
	r := gin.Default()
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", controllers.Register)
		userGroup.POST("/login", controllers.Login)
	}
	listGroup := r.Group("/list")
	{
		listGroup.POST("/create", controllers.CreateList)
		listGroup.GET("/get", controllers.GetLists)
	}
	taskGroup := r.Group("/task")
	{
		taskGroup.POST("/add", controllers.AddTask)
	}
	return r
}
