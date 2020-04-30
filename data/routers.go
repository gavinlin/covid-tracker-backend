package data

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DataRegister(router *gin.RouterGroup) {
	router.GET("/id/:country_id", DataListRetrieve)
	router.GET("/latest", LatestDataRetrieve)
	router.GET("/daily", DailyDataRetrieve)
}

func DataListRetrieve(c *gin.Context) {
	countryStr := c.Param("country_id")
	countryID, err := strconv.Atoi(countryStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := GetData(uint(countryID))

	serializer := DataListSerializer{c, data}
	c.JSON(http.StatusOK, gin.H{"data": serializer.Response()})
}

func LatestDataRetrieve(c *gin.Context) {
	latestData, err := GetLatestData()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	serializer := LatestDataSerializer{c, latestData}
	c.JSON(http.StatusOK, gin.H{"data": serializer.Response()})
}

func DailyDataRetrieve(c *gin.Context) {
	allDates, err := GetAllDates()
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Get all dates length ", len(allDates), " and first one ", allDates[0])
	}
}