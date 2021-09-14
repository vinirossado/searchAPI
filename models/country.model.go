package models

type Country struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Population  int         `json:"population"`
	Language    string      `json:"language"`
	Capital     string      `json:"capital"`
	CountryCode string      `json:"countryCode"`
	Currency    Currency    `json:"currency"`
	Iso         string      `json:"iso"`
	Gmt         int         `json:"gmt"`
	Description interface{} `json:"description"`
}
