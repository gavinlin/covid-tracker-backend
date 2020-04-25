package main

import (
	// "database/sql"
	"encoding/csv"
	// "strconv"
	// "strings"
	// "time"

	// "fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	// "time"
	//"github.com/go-co-op/gocron"
	"github.com/gavinlin/covid-tracker-backend/data"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hellow world"))
}

func readCSVFromURL(url string) ([][] string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)
	reader.Comma = ','
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func task(channel chan[][]string) {
	const url = "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_confirmed_global.csv"

	data, err := readCSVFromURL(url)
	if err != nil {
		log.Println(err)
	}
	channel <- data
}

func main() {
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", home)

	// s1 := gocron.NewScheduler(time.UTC)
	// s1.Every(10).Second().Do(task)
	// s1.Start()

	// log.Println("Start server on :5000")
	// err := http.ListenAndServe(":5000", mux)
	// log.Fatal(err)

	channel := make(chan [][]string)
	go task(channel)
	csvdata := <- channel
	dataService := data.NewFakeDataService()
	dataService.InitDatabase(csvdata)

	/*
	connStr := "postgres://postgres:apple@localhost/covid-19?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	
	db.Query("DELETE FROM country")
	sqlStatement := `INSERT INTO country (country, state, lat, long) VALUES ($1, $2, $3, $4) RETURNING id`
	dataSQLStatement := `INSERT INTO data (date, country_id, confirmed) VALUES ($1, $2, $3)`
	for i, s := range data {
		if i != 0 {
			stmt, err := db.Prepare(sqlStatement)
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
					_ , err = db.Exec(dataSQLStatement, date, countryID, confirmed)
					if err != nil {
						panic(err)
					}
				}
			}
		}
	}

	rows, err := db.Query("SELECT * FROM country")
	if err != nil {
		log.Println(err)
	}
	log.Println("has next row", rows.Next())
	*/
}