package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// Connect initializes a connection to the MySQL database using Gorm.
func Connect() {
	dsn := "root:anish@tcp(127.0.0.1:3306)/anish?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	
	db = d
}

// GetDB returns the initialized Gorm database instance.
func GetDB() *gorm.DB {
	return db
}
