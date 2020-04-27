package main

import (
	"encoding/csv"
	"time"

	"log"
	"net/http"

	"github.com/gavinlin/covid-tracker-backend/common"
	"github.com/gavinlin/covid-tracker-backend/data"
	"github.com/gavinlin/covid-tracker-backend/countries"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
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

func startDataScheduler () {
	s1 := gocron.NewScheduler(time.UTC)
	s1.Every(5).Hours().Do(task)
	s1.Start()
}

func main() {

	db := common.Init()
	defer db.Close()
	dataService := data.NewDBDataService(db)

	mainStruct = MainStruct{
		DataService: dataService,
	}

	startDataScheduler()

	// mux := http.NewServeMux()
	// mux.HandleFunc("/", home)
	// log.Println("Start server on :5000")
	// err := http.ListenAndServe(":5000", mux)

	r := gin.Default()

	v1 := r.Group("/api")

	countries.CountriesRegister(v1.Group("/countries"))
	r.Run()
}