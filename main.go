package main

import (
	"encoding/csv"

	"log"
	"net/http"
	_ "github.com/lib/pq"

	//"github.com/go-co-op/gocron"
	"github.com/gavinlin/covid-tracker-backend/data"
	"github.com/gavinlin/covid-tracker-backend/common"
)

type MainStruct struct {
	DataService data.DataService
}

var mainStruct MainStruct

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

func downloadData(channel chan[][]string) {
	const url = "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_confirmed_global.csv"

	data, err := readCSVFromURL(url)
	if err != nil {
		log.Println(err)
	}
	channel <- data
}

func task() {
	channel := make(chan [][]string)
	go downloadData(channel)
	csvdata := <- channel

	mainStruct.DataService.UpdateDatabase(csvdata)
}

// func initDB() *sql.DB {
// 	connStr := "postgres://postgres:apple@localhost/covid-19?sslmode=disable"
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return db
// }

func startDataScheduler () {
	// s1 := gocron.NewScheduler(time.UTC)
	// s1.Every(10).Second().Do(task)
	// s1.Start()
	task()
}

func main() {

	db := common.Init()
	dataService := data.NewPostgresDataService(db)

	mainStruct = MainStruct{
		DataService: dataService,
	}

	startDataScheduler()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Start server on :5000")
	err := http.ListenAndServe(":5000", mux)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
}