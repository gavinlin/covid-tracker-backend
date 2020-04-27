package services

import (
	"log"
	"strconv"
	"strings"
	"time"
	"github.com/jinzhu/gorm"

	"github.com/gavinlin/covid-tracker-backend/data"
	"github.com/gavinlin/covid-tracker-backend/countries"

)

type dBDataService struct {
	DB *gorm.DB
}

func NewDBDataService(db *gorm.DB) DataService {
	return &dBDataService{
		DB: db,
	}
}

func (p *dBDataService) InitDatabase(incomingData [][]string) {
		p.DB.DropTableIfExists(&countries.Country{})
		p.DB.DropTableIfExists(&data.Data{})
		p.DB.AutoMigrate(&countries.Country{})
		p.DB.AutoMigrate(&data.Data{})

	 	for i, s := range incomingData{
			if i != 0 {
				country := countries.Country{
					Country: s[1],
					State: s[0],
					Lat: s[2],
					Long: s[3],
				}
				p.DB.Create(&country)
				log.Println("country id is ", country.ID)
				for i, confirmed := range s {
					if (i > 3) {
						currentDate := getTime(incomingData[0][i])
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
}

func getTime(original string) time.Time {
	dateSlice := strings.Split(original, "/")
	year, _:= strconv.Atoi("20" + dateSlice[2])
	month, _ := strconv.Atoi(dateSlice[0])
	day, _ := strconv.Atoi(dateSlice[1])
	
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return date
}

func (p *dBDataService) UpdateDatabase(incomingData [][]string) {
	for i, s := range incomingData{
		if i != 0 {
			country := countries.Country{
				Country: s[1],
				State: s[0],
				Lat: s[2],
				Long: s[3],
			}
			p.DB.Where(country).FirstOrCreate(&country)
			
			log.Println("country id is ", country.ID)
			for i, confirmed := range s {
				if (i > 3) {
					currentDate := getTime(incomingData[0][i])
					confirmedNum, _ := strconv.Atoi(confirmed)
					d := data.Data {
						Date: currentDate,
						Confirmed: confirmedNum,
						CountryID: country.ID,
					}
					p.DB.Where(&d).FirstOrCreate(&d)
				}
			}
		}
	}
}