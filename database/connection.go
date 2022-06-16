package database

import (
	"github.com/1deep1/deepcraft-backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open(key), &gorm.Config{})

	if err != nil {
		panic("No connection to database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
