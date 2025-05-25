package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MW(base, login string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Cookie(CookieKey)
		if err != nil {
			ctx.Abort()
			ctx.Redirect(http.StatusFound, login)
			return
		}

		if err := Check(cookie); err != nil {
			ctx.Abort()
			ctx.SetCookie(CookieKey, "", 0, base, "", false, false)
			ctx.Redirect(http.StatusFound, login)
			return
		}

		ctx.Next()
	}
}
