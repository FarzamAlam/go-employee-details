package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {
	e := godotenv.Load() // load .env file
	if e != nil {
		fmt.Println(e)
	}
	username := os.Getenv("db_username")
	password := os.Getenv("db_password")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")
	dbType := os.Getenv("db_type")
	dbName := os.Getenv("db_name")

	dbUri := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", username, password, dbHost, string(dbPort), string(dbName))
	fmt.Println(dbUri)

	conn, err := gorm.Open(dbType, dbUri)
	if err != nil {
		fmt.Println("Error while connecting to database", err)
	}
	db = conn
	db.Debug().AutoMigrate(&Employee{})
	fmt.Println("Connection Successful!")
}

// return a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
