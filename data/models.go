package data

import (
	"fmt"
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

type LatestData struct {
	Date time.Time `json:"date"`
	Confirmed int `json:"confirmed"`
	Recovered int `json:"recovered"`
	Death int `json:"death"`
}

func GetData(countryID uint) ([]Data, error) {
	db := common.GetDB()
	var data = []Data{}

	err := db.Where("country_id = ?", countryID).Find(&data).Error

	return data, err
}

type SumResult struct{
	N int
}

func GetLatestData() (LatestData, error) {
	db := common.GetDB()
	var confirmedResult SumResult
	var deathResult SumResult
	var recoveredResult SumResult
	var data Data
	err := db.Table("data").Order("date DESC").Limit(1).Find(&data).Error
	if err != nil {
		fmt.Println(err)
		return LatestData{}, err
	}

	err = db.Table("data").Select("sum(confirmed) as n").Where("date = ?", data.Date).Scan(&confirmedResult).Error
	if err != nil {
		fmt.Println(err)
		return LatestData{}, err
	}
	err = db.Table("data").Select("sum(death) as n").Where("date = ?", data.Date).Scan(&deathResult).Error
	if err != nil {
		fmt.Println(err)
		return LatestData{}, err
	}
	err = db.Table("data").Select("sum(recovered) as n").Where("date = ?", data.Date).Scan(&recoveredResult).Error
	if err != nil {
		fmt.Println(err)
		return LatestData{}, err
	}

	latestData := LatestData {
		Date: data.Date,
		Confirmed: confirmedResult.N,
		Death: deathResult.N,
		Recovered: recoveredResult.N,
	}
	return latestData, nil
}