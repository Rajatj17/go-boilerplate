package main

import (
	"github.com/gin-gonic/gin"
	"main.go/controller"
	"main.go/library"
	"main.go/models"
)

func main() {
	r := gin.Default()

	r.Use(gin.Logger())

	models.ConnectToDb()

	library.CreateQueue()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"success": true,
			"message": "Your Server is Up!",
		})
	})

	v1 := r.Group("/api/v1")
	{
		v1.POST("user", controller.UserController{}.Create)
	}

	r.Run()
}
