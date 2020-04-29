package data

import (
	"time"

	"github.com/gavinlin/covid-tracker-backend/common"
)

type Data struct {
	ID uint `gorm:"primary_key"`
	Date time.Time `gorm:"column:date"`
	Confirmed int `gorm:"column:confirmed"`
	Recovered int `gorm:"column:recovered"`
	Death int `gorm:"column:death"`
	CountryID uint
}

func GetData(countryID uint) ([]Data, error) {
	db := common.GetDB()
	var data = []Data{}

	err := db.Where("country_id = ?", countryID).Find(&data).Error

	return data, err
}