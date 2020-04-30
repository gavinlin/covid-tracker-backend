package services

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/gavinlin/covid-tracker-backend/countries"
	"github.com/gavinlin/covid-tracker-backend/data"
)

type dBDataService struct {
	DB *gorm.DB
}

func NewDBDataService(db *gorm.DB) DataService {
	return &dBDataService{
		DB: db,
	}
}

func (p *dBDataService) InitDatabase(confirmedData [][] string, recoveredData [][]string, deathData [][]string) {
	p.DB.DropTableIfExists(&countries.Country{})
	p.DB.DropTableIfExists(&data.Data{})
	p.DB.AutoMigrate(&countries.Country{})
	p.DB.AutoMigrate(&data.Data{})

	m := make(map[string]countries.Country)

	for i, s := range confirmedData{
		if i != 0 {
			country := getCountry(s, p.DB)
			countryKey := getCountryKey(s);
			m[countryKey] = country
			for j, confirmed := range s {
				if (j > 3) {
					currentDate := getTime(confirmedData[0][j])
					confirmedNum, _ := strconv.Atoi(confirmed)
					d := data.Data {
						Date: currentDate,
						Confirmed: confirmedNum,
						CountryID: country.ID,
					}
					p.DB.Create(&d)
				}
			}
		}
	}

	udpateDeathData(deathData, p.DB, m)
	updateRecoveredData(recoveredData, p.DB, m)
}

func getCountryKey(row []string) (string) {
	return fmt.Sprintf("%s+%s",row[1], row[0]);
}

func getTime(original string) time.Time {
	dateSlice := strings.Split(original, "/")
	year, _:= strconv.Atoi("20" + dateSlice[2])
	month, _ := strconv.Atoi(dateSlice[0])
	day, _ := strconv.Atoi(dateSlice[1])
	
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return date
}

func (p *dBDataService) UpdateDatabase(confirmedData [][] string, recoveredData [][]string, deathData [][]string) {
	m := make(map[string]countries.Country)
	for i, s := range confirmedData{
		if i != 0 {
			country := getCountry(s, p.DB)
			m[getCountryKey(s)] = country
			log.Println("update country id is ", country.ID)
			for j, confirmed := range s {
				if (j > 3) {
					currentDate := getTime(confirmedData[0][j])
					confirmedNum, _ := strconv.Atoi(confirmed)
					d := data.Data {
						Date: currentDate,
						Confirmed: confirmedNum,
						CountryID: country.ID,
					}
					updateData(currentDate, country.ID, d, p.DB)
				}
			}
		}
	}
	udpateDeathData(deathData, p.DB, m)
	updateRecoveredData(recoveredData, p.DB, m)
}

func udpateDeathData(deathData [][]string, db *gorm.DB, m map[string]countries.Country) {
	for i, s:= range deathData {
		if i != 0 {
			// country := getCountry(s, db)
			country := m[getCountryKey(s)]
			for j, death := range s {
				if (j > 3) {
					currentDate := getTime(deathData[0][j])
					deathNum, _ := strconv.Atoi(death)
					d := data.Data {
						Death: deathNum,
					}
					updateData(currentDate, country.ID, d, db)
				}
			}
		}
	}
}

func updateRecoveredData(recoveredData [][]string, db *gorm.DB, m map[string]countries.Country) {
	for i, s:= range recoveredData {
		if i != 0 {
			// country := getCountry(s, db)
			country := m[getCountryKey(s)]
			for j, death := range s {
				if (j > 3) {
					currentDate := getTime(recoveredData[0][j])
					recoveredNum, _ := strconv.Atoi(death)
					d := data.Data {
						Recovered: recoveredNum,
					}
					updateData(currentDate, country.ID, d, db)
				}
			}
		}
	}
}

func updateData(date time.Time, countryID uint, d data.Data, db *gorm.DB) {
	var updatedData data.Data
	db.Where(data.Data{Date: date, CountryID: countryID}).Assign(&d).FirstOrCreate(&updatedData)
}

func getCountry(array [] string, db *gorm.DB) (countries.Country) {
			floatLat, _ := strconv.ParseFloat(array[2], 64)
			floatLong, _ := strconv.ParseFloat(array[3], 64)
			country := countries.Country{
				Country: array[1],
				State: array[0],
				Lat: strconv.FormatFloat(floatLat, 'f', 2, 64),
				Long: strconv.FormatFloat(floatLong, 'f', 2, 64),
			}
			db.Where(country).FirstOrCreate(&country)
			return country
}