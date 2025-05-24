package setup

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/melsincostan/borders/auth/key"
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

type setupDTO struct {
	Username             string `form:"username" binding:"required"`
	Password             string `form:"password" binding:"required"`
	PasswordConfirmation string `form:"password-confirm" binding:"required"`
}

var didSetup = false

func notTwice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if didSetup {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "setup already completed successfully - restart the application to make the setup endpoints go away",
			})
		}
	}
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

func show() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.HTML(http.StatusOK, "setup.html", pageInfo{
			Transports: transport.Default,
			Countries:  country.Default,
		})
	}
}

func run(db *gorm.DB, jwtKeyConfig *key.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var params setupDTO

		if err := ctx.ShouldBind(&params); err != nil {
			ctx.Error(err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "missing parameters in the request",
			}) // this should respect the users dark / light setting (a webpage with just text might not).
			return
		}

		if params.Password != params.PasswordConfirmation {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "password doesn't match the confirmation",
			}) // this should respect the users dark / light setting (a webpage with just text might not).
			return
		}

		if len(params.Password) < 4 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "please use a password with more than four characters", // not technically correct if a multi-byte character is used.
			}) // this should respect the users dark / light setting (a webpage with just text might not).
			return
		}

		if len(params.Username) < 1 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "please do not use an empty username",
			}) // this should respect the users dark / light setting (a webpage with just text might not).
			return
		}

		if err := Setup(db, params.Username, params.Password); err != nil {
			ctx.Error(err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "error setting up the application - please check the server log",
			})
			return
		}

		didSetup = true
		ctx.Redirect(http.StatusFound, "/")
	}
}
