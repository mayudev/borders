package admin

import (
	"embed"
	"html/template"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//go:embed templates/*.html
var templates embed.FS

func Template(tmpl *template.Template) (err error) {
	tmpl, err = tmpl.ParseFS(templates, "templates/*.html")
	if err != nil {
		return err
	}
	return
}

func Install(group *gin.RouterGroup, db *gorm.DB) {
	group.GET("/login", loginForm())
	group.POST("/login")
}
