package data

import (
	"time"
)

type Data struct {
	ID uint `gorm:"primary_key"`
	Date time.Time `gorm:"column:date"`
	Confirmed int `gorm:"column:confirmed"`
	CountryID uint
}