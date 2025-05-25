package main

import (
	"errors"
	"html/template"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/melsincostan/borders/auth/key"
	"github.com/melsincostan/borders/pages/admin"
	"github.com/melsincostan/borders/pages/stats"
	"github.com/melsincostan/borders/setup"
	"github.com/melsincostan/borders/static"
	"github.com/melsincostan/borders/utils"
	"github.com/melsincostan/envconfig"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func handleSignals(done chan bool) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT, syscall.SIGABRT)
	s := <-signals
	log.Printf("Received a(n) '%s' signal, quitting...", s.String())
	done <- true
}

func main() {

	done := make(chan bool, 1)

	go handleSignals(done)

	db, err := gorm.Open(sqlite.Open("borders.db"), &gorm.Config{
		Logger: logger.Discard,
	})
	if err != nil {
		log.Fatalf("Error opening the database: %s", err.Error())
	}

	appConfig, err := envconfig.Parse[Config]()
	if err != nil {
		log.Fatalf("Could not load the app configuration")
	}

	jwtKeyConfig, err := envconfig.Parse[key.Config]()
	if err != nil {
		log.Fatalf("Could not load the JWT Key configuration: %s", err)
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(gin.Recovery(), utils.Logging())

	tmpl := template.New("pages")

	if err := stats.Template(tmpl); err != nil {
		log.Fatalf("Error setting up statistics template files: %s", err)
	}

	if err := admin.Template(tmpl); err != nil {
		log.Fatalf("Error setting up administrative template files: %s", err)
	}

	if err := setup.IsSetup(db); errors.Is(err, setup.ErrNotSetup) {
		if err := setup.Templates(tmpl); err != nil {
			log.Fatalf("Error setting up the setup templates files%s", err)
		}

		router.SetHTMLTemplate(tmpl)

		if err := setup.Install(router.Group("/setup"), db, jwtKeyConfig, appConfig.Port); err != nil {
			log.Fatalf("Could not install setup endpoints: %s", err.Error())
		}
	} else if err != nil {
		log.Fatalf("Could not check for setup status: %s", err)
	} else {
		router.SetHTMLTemplate(tmpl)
	}

	if err := key.Load(jwtKeyConfig); err != nil {
		log.Fatalf("Error loading JWT Signing key: %s", err)
	}

	static.Install(router.Group("/static"))
	stats.Install(router.Group("/"), db)
	admin.Install(router.Group("/admin"), db)

	go router.Run(appConfig.Addr())
	log.Printf("Application is now listening on '%s'", appConfig.Addr())
	<-done
	rdb, _ := db.DB()
	rdb.Close()
	log.Printf("goodbye!")
}
