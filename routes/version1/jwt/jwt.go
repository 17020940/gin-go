package jwt

import (
	jwtBussiness "go-gin/business/jwt"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	jwtRouter := router.Group("/jwt")
	{
		jwtRouter.GET("/", jwtBussiness.GenerateJwt)

		jwtRouter.GET("/refesh", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Token is refesh",
			})
		})
	}
}
