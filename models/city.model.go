package models

type City struct {
	Id        int    `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type CreateCity struct {
	Name      string `json:"name" binding:"required"`
	IdCountry int    `json:"id_country" binding:"required"`
	Latitude  string `json:"latitude" binding:"required"`
	Longitude string `json:"longitude" binding:"required"`
}

type UpdateCity struct {
	Name      string `json:"name"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}
