package app

import (
	"fmt"
	"golang-mygram/model/domain"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	root     = goDotEnvVariable("DB_USER")
	password = goDotEnvVariable("DB_PASWORD")
	dbName   = goDotEnvVariable("DB_NAME")
	DB       *gorm.DB
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func StartDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", root, password, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error")
	}

	DB = db

	db.AutoMigrate(&domain.User{}, &domain.Photo{})
	fmt.Println("DB connected!!!")
}

func GetDB() *gorm.DB {
	return DB
}
