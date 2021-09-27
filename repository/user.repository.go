package repository

import (
	"gorm.io/gorm"
	"main.go/models"
)

type UserRepository struct {
	ConnectionString *gorm.DB
}

func (instance UserRepository) Create(payload models.User) models.User {
	db := instance.ConnectionString

	db.Create(&payload)

	return payload
}

func (instance UserRepository) Update(id int, payload models.User) models.User {
	db := instance.ConnectionString

	db.Where("id = ?", id).Updates(payload)

	return payload
}
