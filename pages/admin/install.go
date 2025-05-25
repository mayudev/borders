package admin

import (
	"embed"
	"fmt"
	"html/template"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/melsincostan/borders/auth"
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
	loginpath := fmt.Sprintf("%s/login", strings.TrimSuffix(group.BasePath(), "/"))
	authmw := auth.MW(group.BasePath(), loginpath)
	group.GET("", authmw, main(db))
	group.GET("/login", loginForm())
	group.POST("/login", login(db, group.BasePath()))
	group.GET("/logout", authmw, logout(group.BasePath(), loginpath))
}
