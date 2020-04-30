package data

import (
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
	serializer := DailyDataSerializer{c, latestData}
	c.JSON(http.StatusOK, gin.H{"data": serializer.Response()})
}

func DailyDataRetrieve(c *gin.Context) {
	allDailyDatas, err := GetDailyDatas()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} 	
	serializer := DailyDataListSerializer{c, allDailyDatas}
	c.JSON(http.StatusOK, gin.H{"daily": serializer.Response()})
}