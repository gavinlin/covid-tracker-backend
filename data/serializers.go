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

type DataResponse struct {
	ID uint `json:"-"`
	Confirmed int `json:"confirmed"`
	Date time.Time `json:"date"`
	CountryID uint `json:"country_id"`
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

