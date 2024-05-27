package routes

import (
	"github.com/gin-gonic/gin"
)

func New(app *gin.Engine) {
	worker := new(DirWorker)

	action := app.Group("/action")
	{
		action.POST("/mkdir/*dirname", worker.CreateDir)
		action.POST("/upload/*dirname", worker.UploadFile)
		action.GET("/download/*filepath", worker.DownloadFile)
	}
	app.GET("/f/*dirname", worker.ListFiles)
}
