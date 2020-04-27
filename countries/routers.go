package countries

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CountriesRegister(router *gin.RouterGroup) {
	router.GET("/", CountriesRetrieve)
}

func CountriesRetrieve(c *gin.Context) {
	countries, err := GetCountries()

	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}	

	serializer := CountriesSerializer{c, countries}
	c.JSON(http.StatusOK, gin.H{"countries": serializer.Response()})
}
