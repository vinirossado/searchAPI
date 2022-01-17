package models

type City struct {
	ID        string `json:"id"`
	IDCountry string `json:"idCountry"`
	Name      string `json:"name"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// type CityRepository interface {
// 	GetAllCity() ([]models.City, error)
// 	GetCity(uint64) (*models.City, error)
// }
