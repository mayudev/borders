package stats

import (
	"embed"
	"html/template"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//go:embed templates/*
var templates embed.FS

func Template(tmpl *template.Template) (err error) {
	tmpl, err = tmpl.ParseFS(templates, "templates/*.html")
	return
}

func Install(group *gin.RouterGroup, db *gorm.DB, ogtitle string) {
	group.GET("", show(db, ogtitle))
}
