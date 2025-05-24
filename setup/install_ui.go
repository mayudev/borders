package setup

import (
	"crypto/rand"
	"embed"
	"encoding/base64"
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
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

func Install(group *gin.RouterGroup, db *gorm.DB) (err error) {
	k, err := key()
	if err != nil {
		return err
	}

	log.Printf("Navigate to http://127.0.0.1:8080/setup/%s (or where the application is configured to be available) to finish setting up the application", k)

	group.Use(ensureKey(k))
	group.GET(":key", show(k))
	group.POST(":key", run(db))
	return
}

func key() (s string, e error) {
	buf := make([]byte, 8)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(buf), nil
}
