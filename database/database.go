package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect() {
	DBHOST := os.Getenv("DB_HOST")
	DBUSER := os.Getenv("DB_USER")
	DBPASSWORD := os.Getenv("DB_PASSWORD")
	DBNAME := os.Getenv("DB_NAME")
	DBPORT := os.Getenv("DB_PORT")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DBUSER, DBPASSWORD, DBHOST, DBPORT, DBNAME)

	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}

	log.Println("Connected to Database!")
}

func Initialize() {

	connectionString := fmt.Sprintf("%s:@tcp(%s:%s)/", os.Getenv("DB_USER"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + os.Getenv("DB_NAME"))
	if err != nil {
		panic(err)
	}
	db.Close()
}
