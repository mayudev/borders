package admin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
