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

func (c *Controller) GetCountries(ctx *gin.Context) {
	var country []models.Country

	response, err := http.Get(os.Getenv("BASE_API_URL") + "/country")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(responseData, &country); err != nil {
		panic(err)
	}

	ctx.JSONP(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": country,
	})
}

func (c *Controller) GetCountry(ctx *gin.Context) {
	var country models.Country

	response, err := http.Get(os.Getenv("BASE_API_URL") + "/country/" + ctx.Params.ByName("id"))

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(responseData, &country); err != nil {
		panic(err)
	}

	ctx.JSONP(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": country,
	})
}
