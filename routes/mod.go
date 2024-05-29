package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wh64dev/wfcloud/middleware"
)

func New(app *gin.Engine) {
	worker := new(DirWorker)
	app.Use(middleware.CheckPriv)

	app.GET("/f/*filepath", worker.RawFiles)
	app.GET("/path/*dirname", worker.ListFiles)
	app.POST("/upload/*dirname", worker.UploadFile)
}
