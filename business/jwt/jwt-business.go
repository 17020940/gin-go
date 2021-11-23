package jwtBussiness

import (
	"go-gin/models"

	"github.com/gin-gonic/gin"
)

type Account struct {
	ID       int
	Username string
}

func GenerateJwt(c *gin.Context) {
	var acc Account
	models.DB.Raw("SELECT id, username FROM accounts WHERE id = 1").Scan(&acc)
	c.JSON(200, acc)
}
