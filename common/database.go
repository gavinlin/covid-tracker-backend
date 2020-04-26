package common

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() *gorm.DB {
	// db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=covid-19 password=appl")
	db, err := gorm.Open("postgres", "postgres://postgres:apple@localhost/covid-19?sslmode=disable")
	if err !=nil {
		panic(err)
	}
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}