package data
import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
)

type postgresDataService struct {
	DB *gorm.DB
}

func NewPostgresDataService(db *gorm.DB) DataService {
	return &postgresDataService{
		DB: db,
	}
}

func (p *postgresDataService) InitDatabase(data [][]string) {
	// sqlStatement := `INSERT INTO country (country, state, lat, long) VALUES ($1, $2, $3, $4) RETURNING id`
	// dataSQLStatement := `INSERT INTO data (date, country_id, confirmed) VALUES ($1, $2, $3)`
	// for i, s := range data {
	// 	if i != 0 {
	// 		stmt, err := p.DB.Prepare(sqlStatement)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		defer stmt.Close()
	// 		var countryID int
	// 		 err = stmt.QueryRow(s[1], s[0], s[2], s[3]).Scan(&countryID)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		for i, confirmed := range s {
	// 			if(i > 3) {
	// 				dateSlice := strings.Split(data[0][i], "/")
	// 				year, _:= strconv.Atoi("20" + dateSlice[2])
	// 				month, _ := strconv.Atoi(dateSlice[0])
	// 				day, _ := strconv.Atoi(dateSlice[1])
					
	// 				date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	// 				_ , err = p.DB.Exec(dataSQLStatement, date, countryID, confirmed)
	// 				if err != nil {
	// 					panic(err)
	// 				}
	// 			}
	// 		}
	// 	}
	// }

}

func getTime(original string) time.Time {
	dateSlice := strings.Split(original, "/")
	year, _:= strconv.Atoi("20" + dateSlice[2])
	month, _ := strconv.Atoi(dateSlice[0])
	day, _ := strconv.Atoi(dateSlice[1])
	
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return date
}

func (p *postgresDataService) UpdateDatabase(data [][]string) {
	// for i, d := range data {
	// 	if i != 0 {
	// 		var country Country
	// 		p.DB.Table("country").Where("country = ? AND state = ?", d[1], d[0]).First(&country)
	// 		fmt.Println("id ", country.ID, " and name ", country.Country)
	// 	}
	// }
	var dataArray []Data
	p.DB.Table("data").Joins("left join country on country.id = data.country_id").Where("country = ? AND state = ?", "Australia", "New South Wales").Find(&dataArray)
	for _, d := range dataArray {
		fmt.Println("result id ", d.ID, " date ", d.Date, " confirmed " + d.Confirmed)
	}

}