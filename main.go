package main

import (
	"log"
	"time"

	"github.com/melsincostan/borders/db/country"
	"github.com/melsincostan/borders/db/crossing"
	"github.com/melsincostan/borders/db/models"
	"github.com/melsincostan/borders/db/transport"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("borders.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening the database: %s", err.Error())
	}

	db.AutoMigrate(&models.Country{}, &models.Crossing{}, &models.Transport{}, &models.User{})
	if _, err := country.CreateBatch(db, models.DefaultCountries); err != nil {
		log.Fatalf("countries: %s", err)
	}
	if _, err := transport.CreateBatch(db, models.DefaultTransports); err != nil {
		log.Fatalf("transports: %s", err)
	}

	if _, err := crossing.Create(db, models.CrossingBase{
		When:        time.Now(),
		BorderCheck: true,
		PapersCheck: true,
	}, 1, 1); err != nil {
		log.Fatalf("c1: %s", err)
	}
	if _, err := crossing.Create(db, models.CrossingBase{
		When:        time.Now(),
		BorderCheck: true,
		PapersCheck: true,
	}, 1, 1); err != nil {
		log.Fatalf("c2: %s", err)
	}
}
