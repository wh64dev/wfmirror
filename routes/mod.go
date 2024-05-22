package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func New(app *gin.Engine) {
	// action := app.Group("/action")
	// {
	// 	action.POST("/upload", func(ctx *gin.Context) {
	// 		form, err := ctx.MultipartForm()
	// 		if err != nil {
	// 			return
	// 		}

	// 		var path string
	// 		raw := form.Value["path"]
	// 		files := form.File["upload[]"]
	// 		if raw[0] == "/" {
	// 			path = ""
	// 		}

	// 		for _, file := range files {
	// 			log.Printf("received upload file: /%s/%s\n", path, file.Filename)
	// 			ctx.SaveUploadedFile(file, fmt.Sprintf("./data%s/%s", path, file.Filename))
	// 		}

	// 		ctx.JSON(200, gin.H{
	// 			"ok":      1,
	// 			"status":  200,
	// 			"count":   len(files),
	// 			"workdir": fmt.Sprintf("/%s", path),
	// 		})
	// 	})
	// }

	app.GET("/", func(ctx *gin.Context) {
		DirWorker(ctx, "")
	})

	app.GET("/:path/*child", func(ctx *gin.Context) {
		origin := ctx.Param("path")
		child := ctx.Param("child")
		var path = origin
		if child != "/" {
			path = fmt.Sprintf("%s%s", origin, child)
		}

		DirWorker(ctx, path)
	})
}
