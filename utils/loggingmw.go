package utils

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logging() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		total := time.Now().Sub(start)
		log.Printf("%d %s %s %s \"%s\"", ctx.Writer.Status(), total.String(), ctx.RemoteIP(), ctx.Request.Method, ctx.Request.RequestURI)
	}
}
