package models

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

var (
	server   = "vdiapp.database.windows.net"
	user     = "VDIApplication"
	password = "VDI!@!12vdi"
	database = "vdi"
)

func ConnectDatabase() {

	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s",
		server, user, password, database)

	db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic("Failed to open database")
	}

	DB = db
}
