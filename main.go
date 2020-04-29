package main

import (
	"encoding/csv"
	"os"
	"time"

	"log"
	"net/http"

	"github.com/gavinlin/covid-tracker-backend/common"
	"github.com/gavinlin/covid-tracker-backend/countries"
	"github.com/gavinlin/covid-tracker-backend/data"
	"github.com/gavinlin/covid-tracker-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	cors "github.com/rs/cors/wrapper/gin"
)

type MainStruct struct {
	DataService services.DataService
}

var mainStruct MainStruct

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hellow world"))
}

func readCSVFromURL(url string) ([][]string, error) {
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

const confirmed_url = "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_confirmed_global.csv"
const recovered_url = "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_recovered_global.csv"
const death_url = "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_deaths_global.csv"

func downloadData(url string, channel chan [][]string) {

	data, err := readCSVFromURL(url)
	if err != nil {
		log.Println(err)
	}
	channel <- data
}

func task() {
	confirmedChannel := make(chan [][]string)
	go downloadData(confirmed_url, confirmedChannel)

	recoveredChannel := make(chan [][]string)
	go downloadData(recovered_url, recoveredChannel)

	deathChannel := make(chan [][]string)
	go downloadData(death_url, deathChannel)

	confirmedData := <-confirmedChannel
	deathData := <- deathChannel
	recoveredData := <- recoveredChannel
	mainStruct.DataService.UpdateDatabase(confirmedData, recoveredData, deathData)
}

func startDataScheduler() {
	s1 := gocron.NewScheduler(time.UTC)
	s1.Every(5).Hours().Do(task)
	s1.Start()
}

func initDatabaseIfNotExist() {
	log.Println("Start Init database")
	fileInfo, err := os.Stat("covid.db")
	if os.IsNotExist(err) || fileInfo.Size() == 0{
		confirmedChannel := make(chan [][]string)
		go downloadData(confirmed_url, confirmedChannel)

		recoveredChannel := make(chan [][]string)
		go downloadData(recovered_url, recoveredChannel)

		deathChannel := make(chan [][]string)
		go downloadData(death_url, deathChannel)

		confirmedData := <-confirmedChannel
		deathData := <- deathChannel
		recoveredData := <- recoveredChannel
		mainStruct.DataService.InitDatabase(confirmedData, recoveredData, deathData)
	}
}

func main() {

	db := common.Init()
	defer db.Close()
	dataService := services.NewDBDataService(db)

	mainStruct = MainStruct{
		DataService: dataService,
	}

	startDataScheduler()

	// mux := http.NewServeMux()
	// mux.HandleFunc("/", home)
	// log.Println("Start server on :5000")
	// err := http.ListenAndServe(":5000", mux)

	initDatabaseIfNotExist()

	r := gin.Default()
  r.Use(cors.Default())

	v1 := r.Group("/api")

	countries.CountriesRegister(v1.Group("/countries"))
	data.DataRegister(v1.Group("/data"))

	r.Run()
}
