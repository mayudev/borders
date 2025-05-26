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

const (
	crossingSubmit = "/crossing"
)

func Install(group *gin.RouterGroup, db *gorm.DB) {
	loginpath := fmt.Sprintf("%s/login", strings.TrimSuffix(group.BasePath(), "/"))
	authmw := auth.Auth(group.BasePath(), loginpath)
	refreshmw := auth.Refresh(group.BasePath())
	group.GET("", authmw, refreshmw, main(db, group.BasePath()))
	group.GET("/login", loginForm())
	group.POST("/login", login(db, group.BasePath()))
	group.GET("/logout", authmw, logout(group.BasePath()))
	group.POST(crossingSubmit, authmw, refreshmw, createCrossing(db, group.BasePath()))
	group.POST("/country/:id/edit", authmw, refreshmw, edit(db, group.BasePath(), "country"))
	group.POST("/transport/:id/edit", authmw, refreshmw, edit(db, group.BasePath(), "transport"))
	group.POST("/country/:id/delete", authmw, refreshmw, delete(db, group.BasePath(), "country"))
	group.POST("/transport/:id/delete", authmw, refreshmw, delete(db, group.BasePath(), "transport"))
	group.POST("/country/create", authmw, refreshmw, create(db, group.BasePath(), "country"))
	group.POST("/transport/create", authmw, refreshmw, create(db, group.BasePath(), "transport"))
	group.POST("/change-password", authmw, refreshmw, changePW(db, group.BasePath()))
}
