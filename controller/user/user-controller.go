package userController

import (
	"go-gin/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	var user models.User
	c.ShouldBindJSON(&user)
	hashPassword, _ := hashPassword(user.Password)
	user.Password = hashPassword
	user.Email = "test"
	user.PhoneNumber = "test"
	models.DB.Create(&user)
	c.JSON(200, user)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
