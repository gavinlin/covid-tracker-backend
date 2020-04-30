package data

import (
	"time"

	"github.com/gin-gonic/gin"
)

type DataSerializer struct {
	C *gin.Context
	Data
}

type DataListSerializer struct {
	C *gin.Context
	DataList []Data
}

type DailyDataSerializer struct {
	C *gin.Context
	DailyData
}

type DailyDataListSerializer struct {
	C *gin.Context
	DailyDataList []DailyData
}

type DataResponse struct {
	ID uint `json:"-"`
	Confirmed int `json:"confirmed"`
	Date time.Time `json:"date"`
	CountryID uint `json:"country_id"`
	Recovered int `json:"recovered"`
	Death int `json:"death"`
}

type DailyDataResponse struct {
	Date time.Time `json:"date"`
	Confirmed int `json:"confirmed"`
	Recovered int `json:"recovered"`
	Death int `json:"death"`
}

func (s *DataSerializer) Response() DataResponse {
	response := DataResponse {
		ID: s.Data.ID,
		Confirmed: s.Data.Confirmed,
		Date: s.Data.Date,
		Recovered: s.Data.Recovered,
		Death: s.Data.Death,
		CountryID: s.Data.CountryID,
	}
	return response
}

func (s *DataListSerializer) Response() []DataResponse {
	response := []DataResponse{}
	for _, data := range s.DataList {
		serializer := DataSerializer{s.C, data}
		response = append(response, serializer.Response())
	}
	return response
}

func (s *DailyDataSerializer) Response() DailyDataResponse {
	response := DailyDataResponse {
		Date: s.DailyData.Date,
		Confirmed: s.DailyData.Confirmed,
		Death: s.DailyData.Death,
		Recovered: s.DailyData.Recovered,
	}
	return response
}

func (s *DailyDataListSerializer) Response() []DailyDataResponse {
	response := []DailyDataResponse{}
	for _, dailyData := range s.DailyDataList {
		serializer := DailyDataSerializer{s.C, dailyData}
		response = append(response, serializer.Response())
	}
	return response
}