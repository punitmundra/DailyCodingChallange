package model

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	fmt.Println("-> [username:", username, "] [password:", password, "] [dbName:", dbName, "] [dbHost:", dbHost, "]")

	var err error
	db, err = gorm.Open("mysql", "root:root@/sqadcast?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("error while connecting to db:", err)
		return
	}

	db.Debug().AutoMigrate(&Incident{}) //Database migration
}

//GetDB ... returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
