package auth

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	expiryContextKey = "jwt-expiry-time"
	refreshWhen      = (Validity / 4) // refresh at the earliest when the token is at 3/4 of the validity period
)

func Auth(base, login string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Cookie(CookieKey)
		if err != nil {
			ctx.Abort()
			ctx.Redirect(http.StatusFound, login)
			return
		}

		if exp, err := Check(cookie); err != nil {
			ctx.Abort()
			ctx.SetCookie(CookieKey, "", 0, base, "", false, false)
			ctx.Redirect(http.StatusFound, login)
			return
		} else {
			ctx.Set(expiryContextKey, exp)
		}

		ctx.Next()
	}
}

func Refresh(base string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer ctx.Next()
		rexp, ok := ctx.Get(expiryContextKey)
		if !ok {
			log.Printf("expiry key not set in context")
			return
		}

		exp, ok := rexp.(time.Time)
		if !ok {
			log.Printf("expiry in context is not a time value")
			return
		}

		if time.Now().After(exp.Add(-refreshWhen)) {
			token, err := Create()
			if err != nil {
				log.Printf("wanted to refresh token but got an error: %s", err)
				return
			}
			ctx.SetCookie(CookieKey, token, int(Validity.Seconds()), base, "", false, true)
			log.Print("refreshed cookie")
		}
	}
}
