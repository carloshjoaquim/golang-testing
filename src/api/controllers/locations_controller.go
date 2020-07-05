package controllers

import (
	"github.com/carloshjoaquim/golang-testing/src/api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCountry(c *gin.Context) {
	country, err := services.LocationService.GetCountry(c.Param("country_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, country)
}
