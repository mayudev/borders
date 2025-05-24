package main

import (
	"errors"
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/melsincostan/borders/setup"
	"github.com/melsincostan/borders/static"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	db, err := gorm.Open(sqlite.Open("borders.db"), &gorm.Config{
		Logger: logger.Default,
	})
	if err != nil {
		log.Fatalf("Error opening the database: %s", err.Error())
	}

	router := gin.Default()

	static.Install(router.Group("/static"))

	tmpl := template.New("pages")

	if err := setup.Templates(tmpl); err != nil {
		log.Fatalf("%s", err)
	}

	router.SetHTMLTemplate(tmpl)

	if err := setup.IsSetup(db); errors.Is(err, setup.ErrNotSetup) {
		if err := setup.Install(router.Group("/setup"), db); err != nil {
			log.Fatalf("Could not install setup endpoints: %s", err.Error())
		}
	} else if err != nil {
		log.Fatalf("Could not check for setup status: %s", err)
	}

	router.Run(":8080")
}
