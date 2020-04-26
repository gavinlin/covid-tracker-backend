package data
import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

type postgresDataService struct {
	DB *sql.DB
}

func NewPostgresDataService(db *sql.DB) DataService {
	return &postgresDataService{
		DB: db,
	}
}

func (p *postgresDataService) InitDatabase(data [][]string) {
	sqlStatement := `INSERT INTO country (country, state, lat, long) VALUES ($1, $2, $3, $4) RETURNING id`
	dataSQLStatement := `INSERT INTO data (date, country_id, confirmed) VALUES ($1, $2, $3)`
	for i, s := range data {
		if i != 0 {
			stmt, err := p.DB.Prepare(sqlStatement)
			if err != nil {
				panic(err)
			}
			defer stmt.Close()
			var countryID int
			 err = stmt.QueryRow(s[1], s[0], s[2], s[3]).Scan(&countryID)
			if err != nil {
				panic(err)
			}
			for i, confirmed := range s {
				if(i > 3) {
					dateSlice := strings.Split(data[0][i], "/")
					year, _:= strconv.Atoi("20" + dateSlice[2])
					month, _ := strconv.Atoi(dateSlice[0])
					day, _ := strconv.Atoi(dateSlice[1])
					
					date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
					_ , err = p.DB.Exec(dataSQLStatement, date, countryID, confirmed)
					if err != nil {
						panic(err)
					}
				}
			}
		}
	}
}

type Country struct {
	ID int `TbField:id`
	Country string `TbField:country`
	State sql.NullString `TbField:state`
	lat float32 `TbField: lat`
	long float32 `TbField: long`
}

func (p *postgresDataService) UpdateDatabase(data [][]string) {
	fmt.Println("Update data to database")
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 1*time.Second)

	countrys := []Country{}

	countryAndRegionStatement := `SELECT * FROM country WHERE country=$1 AND state=$2`
	for i, d := range data {
		if i != 0 {
			rows, err := p.DB.QueryContext(ctx, countryAndRegionStatement, d[1], d[0])
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()
			for rows.Next() {
				var country Country
				err = rows.Scan(
					&country.ID,
					&country.Country, 
					&country.State,
					&country.lat,
					&country.long)

				if err != nil {
					log.Fatal(err)
				}
				countrys = append(countrys, country)
			}
		}
	}
	fmt.Println("array size is ", len(countrys))
	for _, c := range countrys {
		fmt.Println("id ", c.ID, " and name ", c.Country )
	}
}