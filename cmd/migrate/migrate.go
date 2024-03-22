package main

import (
	"assignment-movie/common/database"
	"assignment-movie/models"
	"fmt"
)

func main() {

	_, err := database.InitDatabase()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	err = database.DB.AutoMigrate(&models.Movies{})
	if err != nil {
		panic("Error Create Database Customers")
	}

	fmt.Println("SUCCESSFULLY ADD MIGRATION")
}
