package data

import (
	"fmt"
	"time"
	"github.com/jinzhu/gorm"

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

type DailyData struct {
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

func GetLatestData() (DailyData, error) {
	db := common.GetDB()
	var data Data
	err := db.Table("data").Order("date DESC").Limit(1).Find(&data).Error
	if err != nil {
		fmt.Println(err)
		return DailyData{}, err
	}
	return GetDataByDate(data.Date, db)
}

func GetAllDates() ([]time.Time, error) {
	var allDates []Data
	var dates []time.Time
	db := common.GetDB()
	err := db.Table("data").Select("DISTINCT date").Scan(&allDates).Error
	for _, data := range allDates {
		dates = append(dates, data.Date)
	}
	return dates, err
}

func GetDailyDatas() ([]DailyData, error) {
	allDailyDatas := []DailyData{}
	allDates, err := GetAllDates()

	if err != nil {
		return allDailyDatas, err
	}

	db := common.GetDB()
	
	for _, date := range allDates {
		dailyData, err := GetDataByDate(date, db)
		if err != nil {
			return allDailyDatas, err
		}
		allDailyDatas = append(allDailyDatas, dailyData)
	}
	return allDailyDatas, nil
}

func GetDataByDate(date time.Time, db *gorm.DB) (DailyData, error) {
	var confirmedResult SumResult
	var deathResult SumResult
	var recoveredResult SumResult
	err := db.Table("data").Select("sum(confirmed) as n").Where("date = ?", date).Scan(&confirmedResult).Error
	if err != nil {
		fmt.Println(err)
		return DailyData{}, err
	}
	err = db.Table("data").Select("sum(death) as n").Where("date = ?", date).Scan(&deathResult).Error
	if err != nil {
		fmt.Println(err)
		return DailyData{}, err
	}
	err = db.Table("data").Select("sum(recovered) as n").Where("date = ?", date).Scan(&recoveredResult).Error
	if err != nil {
		fmt.Println(err)
		return DailyData{}, err
	}

	latestData := DailyData {
		Date: date,
		Confirmed: confirmedResult.N,
		Death: deathResult.N,
		Recovered: recoveredResult.N,
	}
	return latestData, nil
}