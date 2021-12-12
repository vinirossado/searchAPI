package main

import (
	"log"
	"net/http"
	"searchAPI/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

func main() {
	r := gin.New()

	gin.SetMode(gin.ReleaseMode)

	c := controllers.NewController()

	v1 := r.Group("/api/v1")
	{
		schools := v1.Group("/schools/")
		{
			schools.GET(":id", c.GetSchool)
			schools.GET("", c.GetSchools)
		}

		cities := v1.Group("/cities")
		{
			cities.GET(":id", c.GetCity)
			cities.GET("", c.GetCities)
		}

		countries := v1.Group("/countries")
		{
			countries.GET(":id", c.GetCountry)
			countries.GET("", c.GetCountries)
		}

		home := v1.Group("/home")
		{
			home.GET(":id", c.GetHome)
			home.GET("", c.GetHome)
		}
	}

	http.ListenAndServe(":8080", r)
}
