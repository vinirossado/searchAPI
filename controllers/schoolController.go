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

func (c *Controller) GetSchools(ctx *gin.Context) {
	var schools []models.School

	response, err := http.Get(os.Getenv("BASE_API_URL") + "/schools")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(responseData, &schools); err != nil {
		panic(err)
	}

	ctx.JSONP(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": schools,
	})
}

func (c *Controller) GetSchool(ctx *gin.Context) {
	var school models.School

	response, err := http.Get(os.Getenv("BASE_API_URL") + "/schools/" + ctx.Params.ByName("id"))

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(responseData, &school); err != nil {
		panic(err)
	}

	ctx.JSONP(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": school,
	})
}
