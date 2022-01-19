package models

type City struct {
	ID          string        `json:"id"`
	IDCountry   string        `json:"idCountry"`
	Name        string        `json:"name"`
	Latitude    string        `json:"latitude"`
	Longitude   string        `json:"longitude"`
	ImageUrl    string        `json:"imageUrl"`
	Blurhash    string        `json:"blurhash"`
	Description interface{}   `json:"description"`
	Country     Country       `json:"country"`
	Images      []interface{} `json:"images"`
	Attractions []interface{} `json:"attractions"`
}

// type CityRepository interface {
// 	GetAllCity() ([]models.City, error)
// 	GetCity(uint64) (*models.City, error)
// }
