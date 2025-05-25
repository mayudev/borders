package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/melsincostan/borders/auth"
	"github.com/melsincostan/borders/db/user"
	"github.com/melsincostan/borders/utils"
	"gorm.io/gorm"
)

type loginDTO struct {
	Name     string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func loginForm() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", nil)
	}
}

func login(db *gorm.DB, base string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var params loginDTO

		if err := ctx.ShouldBind(&params); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if err := user.Authenticate(db, params.Name, params.Password); err != nil {
			ctx.Error(err)
			ctx.AbortWithStatusJSON(http.StatusForbidden, utils.ErrJSON("could not authenticate user"))
			return
		}

		token, err := auth.Create()
		if err != nil {
			ctx.Error(err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrJSON("could not create token"))
			return
		}
		ctx.SetCookie(auth.CookieKey, token, int(auth.Validity.Seconds()), base, "", false, true)
		ctx.Redirect(http.StatusFound, base)
	}
}

func logout(base string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.SetCookie(auth.CookieKey, "", 0, base, "", false, false)
		ctx.Redirect(http.StatusFound, "/")
	}
}
