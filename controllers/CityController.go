package controllers

import (
	"net/http"
	"searchAPI/models"

	"github.com/gin-gonic/gin"
)

func FindCities(c *gin.Context) {

	var cities []models.City
	models.DB.Table("City").Find(&cities)

	c.JSON(http.StatusOK, gin.H{"data": cities})

}

func CreateCity(c *gin.Context) {
	var input models.CreateCity

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	city := models.City{Name: input.Name, Longitude: input.Latitude, Latitude: input.Longitude}
	models.DB.Create(&city)

	c.JSON(http.StatusOK, gin.H{"data": city})
}

func FindCity(c *gin.Context) { // Get model if exist
	var city models.City

	if err := models.DB.Where("id = ?", c.Param("id")).First(&city).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": city})
}

func UpdateCity(c *gin.Context) {
	// Get model if exist
	var city models.City
	if err := models.DB.Table("City").Where("id = ?", c.Param("id")).First(&city).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UpdateCity
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&city).Updates((models.City{Name: input.Name}))
	models.DB.Table("City").Save(&city)
	c.JSON(http.StatusOK, gin.H{"data": city})
}
