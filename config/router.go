package config

import (
	"github.com/gin-gonic/gin"
	"github.com/pikejun/mggp/controller/v1"
)

// 新建路由
func NewRouter() *gin.Engine {
	// 新建一个没有任何默认中间件的路由
	r := gin.New()
	// 全局中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	user := v1.NewUserController()
	// 路由组
	apiv1 := r.Group("/api/v1")
	{
		// Restful 路由
		/* apiv1.POST("/users", func(c *gin.Context) {
			c.JSON(200, "add user success!")
		}) */
		apiv1.POST("/users/creatUser", user.Create)
		apiv1.DELETE("/users/deleteUserById/:id", user.Delete)
		apiv1.PUT("/users/updateUserById/:id", user.Update)
		apiv1.GET("/users/getUserById/:id", user.GetUserById)
		apiv1.GET("/users/listUserByPage", user.ListUserByPage)
	}

	apiv2 := r.Group("/api/v2")
	{
		apiv2.POST("/users")
	}


	return r
}
