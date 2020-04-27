package countries

import (
	"github.com/gavinlin/covid-tracker-backend/common"
	"github.com/gavinlin/covid-tracker-backend/data"
)

func GetCountries() ([]data.Country, error) {
	db := common.GetDB()
	var countries []data.Country

	err := db.Find(&countries).Error

	return countries, err
}