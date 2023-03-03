package router

import (
	"db-go-gin/internal/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (*UserRouter) InitUserRouter(userGroup *gin.RouterGroup) {
	userGroup.POST("login", userController.Login)
	userGroup.POST("register", userController.Register)

	// 登录权限验证
	userGroup.Use(middleware.JwtAuth())

	userGroup.POST("logout", userController.Logout)

	// 身份权限验证
	userGroup.Use(middleware.CasbinHandler())
}
