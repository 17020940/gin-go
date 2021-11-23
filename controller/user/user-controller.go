package userController

import (
	"go-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	var user = &models.User{}
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Password, _ = hashPassword(user.Password)
	models.DB.Create(user)
	c.JSON(200, user)
}

func hashPassword(password *string) (*string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	hashPassword := string(bytes)
	return &hashPassword, err
}
