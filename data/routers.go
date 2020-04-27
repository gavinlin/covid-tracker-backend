package data

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DataRegister(router *gin.RouterGroup) {
	router.GET("/:country_id", DataListRetrieve)
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