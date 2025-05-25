package setup

import (
	"crypto/rand"
	"embed"
	"encoding/base64"
	"fmt"
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/melsincostan/borders/auth/key"
	"gorm.io/gorm"
)

//go:embed templates/*
var templates embed.FS

func Templates(tmpl *template.Template) (err error) {
	tmpl, err = tmpl.ParseFS(templates, "templates/*.html")
	if err != nil {
		return err
	}
	return
}

func Install(group *gin.RouterGroup, db *gorm.DB, jwtKeyConfig *key.Config, port uint) (err error) {
	k, err := setupKey()
	if err != nil {
		return err
	}

	if err := key.Create(jwtKeyConfig); err != nil { // setup the JWT key now so it can be loaded during startup, even when doing the setup
		return fmt.Errorf("could not create a JWT key: %w", err)
	}

	log.Printf("Navigate to http://127.0.0.1:%d/setup/%s (or where the application is configured to be available) to finish setting up the application", port, k)

	group.Use(ensureKey(k), notTwice())
	group.GET(":key", show())
	group.POST(":key", run(db))
	return
}

func setupKey() (s string, e error) {
	buf := make([]byte, 8)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(buf), nil
}
