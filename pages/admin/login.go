package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func loginForm() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", nil)
	}
}
