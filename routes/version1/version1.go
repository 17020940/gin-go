package version1

import (
	userRouter "go-gin/routes/version1/user"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	v1 := router.Group("/v1")
	userRouter.Routes(v1)
}
