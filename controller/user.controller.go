package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/models"
	"main.go/repository"
	"main.go/validator"
)

type UserController struct{}

func (user UserController) Create(c *gin.Context) {
	var json validator.CreateUser

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbConnection := models.DbPrimary

	repo := repository.UserRepository{ConnectionString: dbConnection}

	repo.Create(models.User{
		FirstName: json.FirstName,
		LastName:  json.LastName,
		Age:       uint(json.Age),
		Password:  json.Password,
		Username:  json.Username,
	})

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "User is created"})
}
