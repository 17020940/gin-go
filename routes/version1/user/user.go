package userRouter

import (
	userController "go-gin/controller/user"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	userRouter := router.Group("/user")
	{
		userRouter.POST("/", userController.CreateUser)
	}
}
