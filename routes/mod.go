package routes

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wh64dev/wfcloud/config"
	"github.com/wh64dev/wfcloud/middleware"
)

func New(app *gin.Engine) {
	dirWorker := new(DirWorker)
	app.Use(middleware.CORS)

	app.GET("/", func(ctx *gin.Context) {
		start := time.Now()
		cnf := config.Get()
		ctx.JSON(200, gin.H{
			"ok":           1,
			"version":      cnf.Dist.Version,
			"directory":    cnf.Global.DataDir,
			"service_name": cnf.Service.Name,
			"respond_time": fmt.Sprintf("%dms", time.Since(start).Milliseconds()),
		})
	})
	app.GET("/path/*dirname", dirWorker.List)
}
