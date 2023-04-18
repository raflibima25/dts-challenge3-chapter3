package database

import (
	"challenge-3-chapter-3/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "123456"
	dbPort   = "5432"
	dbName   = "role-api"
	db       *gorm.DB
	err      error
)

func StartDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", host, dbPort, user, password, dbName)

	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database:", err)
	}

	log.Println("Success connect database")

	db.Debug().AutoMigrate(models.Role{}, models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	return db
}
