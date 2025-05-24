package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := gorm.Open(sqlite.Open("borders.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening the database: %s", err.Error())
	}

}
