package admin

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/melsincostan/borders/db/country"
	"github.com/melsincostan/borders/db/models"
	"github.com/melsincostan/borders/db/transport"
	"github.com/melsincostan/borders/utils"
	"gorm.io/gorm"
)

type adminObj struct {
	Countries      []models.Country
	Transports     []models.Transport
	CreateCrossing string
	BorderValue    string
	PapersValue    string
}

func main(db *gorm.DB, base string) gin.HandlerFunc {
	createCrossing := fmt.Sprintf("%s%s", strings.TrimSuffix(base, "/"), crossingSubmit)
	return func(ctx *gin.Context) {
		countries, err := country.Read(db)
		if err != nil {
			ctx.Error(err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrJSON("could not load countries"))
			return
		}

		transports, err := transport.Read(db)
		if err != nil {
			ctx.Error(err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrJSON("could not load transports"))
		}

		ctx.HTML(http.StatusOK, "admin.html", adminObj{
			Countries:      countries,
			Transports:     transports,
			CreateCrossing: createCrossing,
			BorderValue:    borderValue,
			PapersValue:    papersValue,
		})
	}
}
