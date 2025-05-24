package main

import (
	"errors"
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/melsincostan/borders/pages/stats"
	"github.com/melsincostan/borders/setup"
	"github.com/melsincostan/borders/static"
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

	router := gin.Default()

	tmpl := template.New("pages")

	if err := stats.Template(tmpl); err != nil {
		log.Fatalf("Error setting up statistics template files: %s", err)
	}

	if err := setup.IsSetup(db); errors.Is(err, setup.ErrNotSetup) {
		if err := setup.Templates(tmpl); err != nil {
			log.Fatalf("Error setting up the setup templates files%s", err)
		}

		router.SetHTMLTemplate(tmpl)

		if err := setup.Install(router.Group("/setup"), db); err != nil {
			log.Fatalf("Could not install setup endpoints: %s", err.Error())
		}
	} else if err != nil {
		log.Fatalf("Could not check for setup status: %s", err)
	} else {
		router.SetHTMLTemplate(tmpl)
	}

	static.Install(router.Group("/static"))
	stats.Install(router.Group("/"), db)

	router.Run(":8080")
}
