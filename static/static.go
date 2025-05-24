package static

import (
	"embed"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

//go:embed files/**/*
var files embed.FS

type fileDTO struct {
	Name string `uri:"file" binding:"required"`
}

func Install(group *gin.RouterGroup) {
	group.GET("*file", serve(files))
}

func serve(files embed.FS) gin.HandlerFunc {
	httpfs := http.FS(files)
	return func(ctx *gin.Context) {
		var params fileDTO

		if err := ctx.ShouldBindUri(&params); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx.FileFromFS(path.Join("files", params.Name), httpfs)
	}
}
