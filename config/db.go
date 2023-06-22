package config

import (
	"gorm.io/driver/postgress"
	"gorm.io/gorm"
)

var DB *gorm.DB
func Connect() {
	db, err := gorm.Open(postgress.Open("postgress://postgress:postgress@localhost:5432/postgress"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Score{})
	DB = db
}