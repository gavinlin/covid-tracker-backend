package countries

import (
	"github.com/gavinlin/covid-tracker-backend/common"
)

type Country struct {
	ID uint `gorm:"primary_key"`
	Country string `gorm:"column:country"`
	State string `gorm:"column:state"`
	Lat string `gorm:"column:lat"`
	Long string `gorm:"column:long"`
}


func GetCountries() ([]Country, error) {
	db := common.GetDB()
	var countries []Country

	err := db.Find(&countries).Error

	return countries, err
}