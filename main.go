package main

import (
	"log"

	"github.com/melsincostan/borders/setup"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	db, err := gorm.Open(sqlite.Open("borders.db"), &gorm.Config{
		Logger: logger.Discard,
	})
	if err != nil {
		log.Fatalf("Error opening the database: %s", err.Error())
	}

	if err := setup.Setup(db, "bwaa", "bwaa"); err != nil {
		log.Fatalf("couldn't setup: %s", err)
	}
}
