package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"searchAPI/models"

	"github.com/gin-gonic/gin"
)

func (c *Controller) GetCities(ctx *gin.Context) {
	var cities []models.City

	response, err := http.Get(os.Getenv("BASE_API_URL") + "/city")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(responseData, &cities); err != nil {
		panic(err)
	}

	ctx.JSONP(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": cities,
	})
}

func (c *Controller) GetCity(ctx *gin.Context) {
	var city models.City

	response, err := http.Get(os.Getenv("BASE_API_URL") + "/city/" + ctx.Params.ByName("id"))

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(responseData, &city); err != nil {
		panic(err)
	}

	ctx.JSONP(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": city,
	})
}
