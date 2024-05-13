package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	m "lgin/models"
)

func Connect() *gorm.DB {

	fmt.Println("Connecting to database...")

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv(("DB_HOST"))
	port := os.Getenv("DB_PORT")
	database_name := os.Getenv("DB_NAME")

	cs := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database_name)
	db, err := gorm.Open(mysql.Open(cs), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
	}

	db.AutoMigrate(&m.Example{})

	return db
}
