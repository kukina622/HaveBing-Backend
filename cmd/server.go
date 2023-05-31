package main

import (
	"fmt"
	"log"
	"os"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
	)
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	r.Run(os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT"))
}
