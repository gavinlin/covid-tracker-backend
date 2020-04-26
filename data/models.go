package data

import (
	"time"
)

type Country struct {
	ID uint `gorm:"primary_key"`
	Country string `gorm:"column:country"`
	State *string `gorm:"column:state"`
	lat float32 `gorm:"column:lat"`
	long float32 `gorm:"column:long"`
}

type Data struct {
	ID uint `gorm:"primary_key"`
	Date time.Time `gorm:"column:date"`
	Confirmed string `gorm:"column:confirmed"`
	CountryID uint
}