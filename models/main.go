package models

import (
	"fmt"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
	"main.go/config"
)

var DbPrimary *gorm.DB

func ConnectToDb() {
	dbConfig := config.AppConfig.DbConfig
	dsn := CreateConnnectionString(dbConfig)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB Connection Error")
	}

	db.AutoMigrate(&User{})

	DbPrimary = db
}

func CreateConnnectionString(dbConfig config.DbConfig) string {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		dbConfig.DbHost,
		dbConfig.DbUser,
		dbConfig.DbPassword,
		dbConfig.DbName,
		dbConfig.DbPort,
	)

	return connectionString
}
