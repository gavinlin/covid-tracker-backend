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

type LatestDataSerializer struct {
	C *gin.Context
	LatestData 
}

type DataResponse struct {
	ID uint `json:"-"`
	Confirmed int `json:"confirmed"`
	Date time.Time `json:"date"`
	CountryID uint `json:"country_id"`
	Recovered int `json:"recovered"`
	Death int `json:"death"`
}

type LatestDataResponse struct {
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

func (s *LatestDataSerializer) Response() LatestDataResponse{
	response := LatestDataResponse {
		Date: s.LatestData.Date,
		Confirmed: s.LatestData.Confirmed,
		Death: s.LatestData.Death,
		Recovered: s.LatestData.Recovered,
	}
	return response
}