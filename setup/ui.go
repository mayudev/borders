package setup

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/melsincostan/borders/db/country"
	"github.com/melsincostan/borders/db/models"
	"github.com/melsincostan/borders/db/transport"
	"gorm.io/gorm"
)

type keyDTO struct {
	Key string `uri:"key" binding:"required"`
}

type pageInfo struct {
	Key        string
	Transports []models.TransportBase
	Countries  []models.CountryBase
}

func ensureKey(key string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var params keyDTO

		if err := ctx.ShouldBindUri(&params); err != nil {
			ctx.AbortWithError(http.StatusNotFound, fmt.Errorf("could not verify key: %w", err))
			return
		}

		if params.Key != key {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		ctx.Next()
	}
}

func show(key string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.HTML(http.StatusOK, "setup.html", pageInfo{
			Key:        key,
			Transports: transport.Default,
			Countries:  country.Default,
		})
	}
}

func run(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := Setup(db, "bwaa", "bwaa"); err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
}
