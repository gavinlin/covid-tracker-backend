package countries

import (
	"github.com/gin-gonic/gin"
)

type CountriesSerializer struct {
	C *gin.Context
	Countries []Country
}

type CountrySerializer struct {
	C *gin.Context
	Country Country
}

type CountryResponse struct {
	ID uint `json:"-"`
	Country string `json:"country"`
	State string `json:"state"`
	Lat string `json:"lat"`
	Long string `json:"long"`
}

func (s *CountrySerializer) Response() CountryResponse {
	response := CountryResponse {
		ID: s.Country.ID,
		Country: s.Country.Country,
		State: s.Country.State,
		Lat: s.Country.Lat,
		Long: s.Country.Long,
	}
	return response
}

func (s *CountriesSerializer) Response() []CountryResponse {
	response := []CountryResponse{}
	for _, country := range s.Countries {
		serializer := CountrySerializer{s.C, country}
		response = append(response, serializer.Response())
	}
	return response
}