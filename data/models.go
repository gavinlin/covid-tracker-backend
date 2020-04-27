package data

import (
	"time"
)

type Country struct {
	ID uint `gorm:"primary_key"`
	Country string `gorm:"column:country"`
	State string `gorm:"column:state"`
	Lat string `gorm:"column:lat"`
	Long string `gorm:"column:long"`
}

type Data struct {
	ID uint `gorm:"primary_key"`
	Date time.Time `gorm:"column:date"`
	Confirmed int `gorm:"column:confirmed"`
	CountryID uint
}