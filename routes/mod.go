package routes

import (
	"github.com/gin-gonic/gin"
)

func New(app *gin.Engine) {
	worker := new(DirWorker)

	app.POST("/mkdir/*dirname", worker.CreateDir)
	app.POST("/upload/*dirname", worker.UploadFile)
	app.GET("/download/*filepath", worker.DownloadFile)
	app.GET("/files/*dirname", worker.ListFiles)
}
