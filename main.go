package main

import (
	"net/http"
	"os"
	"searchAPI/controllers"
	"searchAPI/models"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
	// })

	
	models.ConnectDatabase()

	r.GET("/cities", controllers.FindCities)
	r.POST("/cities", controllers.CreateCity)
	r.GET("/cities/:id", controllers.FindCity)
	r.PATCH("/cities/:id", controllers.UpdateCity) // new

	x := ":" + os.Getenv("ASPNETCORE_PORT")
	http.ListenAndServe(x, r)

}
