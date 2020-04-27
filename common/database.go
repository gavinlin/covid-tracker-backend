package common

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() *gorm.DB {
	// db, err := gorm.Open("postgres", "postgres://postgres:apple@localhost/covid-19?sslmode=disable")
	db, err := gorm.Open("sqlite3", "./covid.db")
	if err !=nil {
		panic(err)
	}
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}