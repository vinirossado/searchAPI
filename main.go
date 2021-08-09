package main

import (
	"net/http"
	"searchAPI/controllers"
	"searchAPI/models"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	models.ConnectDatabase()

	// response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	// if err != nil {
	// 	fmt.Print(err.Error())
	// 	os.Exit(1)
	// }

	// responseData, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(responseData))

	r.GET("/cities", controllers.FindCities)
	r.POST("/cities", controllers.CreateCity)
	r.GET("/cities/:id", controllers.FindCity)
	r.PATCH("/cities/:id", controllers.UpdateCity)

	http.ListenAndServe(":8080", r)

}
