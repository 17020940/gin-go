package main

import (
	"go-gin/models"
	"go-gin/routes/version1"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	initEnviroment()
	models.ConnectDatabase()
}

func main() {
	r := gin.Default()
	version1.Routes(r)
	r.Run()
}

func initEnviroment() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
