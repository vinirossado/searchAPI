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

func (c *Controller) GetHome(ctx *gin.Context) {
	var destinations []models.Destination

	response, err := http.Get(os.Getenv("BASE_API_URL") + "/home")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(responseData, &destinations); err != nil {
		panic(err)
	}

	ctx.JSONP(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": destinations,
	})
}
