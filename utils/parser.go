package parser

import (
	"encoding/json"
	"searchAPI/models"
)

func byteArrayToJson(bytes []byte, T *models.Cities) {

	if err := json.Unmarshal(bytes, &T); err != nil {
		panic(err)
	}

}
