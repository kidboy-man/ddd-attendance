package database

import (
	"fmt"
	"log"

	"github.com/kidboy-man/ddd-attendance/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var con *gorm.DB

func Setup() (err error) {
	dbUri := fmt.Sprintf(
		"host=%s port=%s, user=%s dbname=%s sslmode=disable password=%s",
		configs.AppConfig.DBHost,
		configs.AppConfig.DBPort,
		configs.AppConfig.DBUser,
		configs.AppConfig.DBName,
		configs.AppConfig.DBPassword,
	)

	log.Println(dbUri)
	db, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
		return
	}

	con = db
	return
}

// returns a handle to the DB object
func GetDB() *gorm.DB {
	return con
}
