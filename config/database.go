package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	var(
		DBhost 		= os.Getenv("DB_HOST")
		DBUser 		= os.Getenv("DB_USER")
		DBPassword 	= os.Getenv("DB_PASSWORD")
		DBName 		= os.Getenv("DB_NAME")
		DBPort 		= os.Getenv("DB_PORT")
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", DBhost, DBUser, DBPassword, DBName, DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("not connect database")
	}

	return db
}