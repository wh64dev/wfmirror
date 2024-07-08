package routes

import (
	"fmt"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wh64dev/wfcloud/config"
	"github.com/wh64dev/wfcloud/middleware"
	"github.com/wh64dev/wfcloud/util"
)

func New(app *gin.Engine) {
	dirWorker := new(DirWorker)
	app.Use(middleware.CORS)
	app.LoadHTMLGlob("pages/*")

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
	app.GET("/raw/*dirname", dirWorker.ShowRaw)
	app.GET("/nodeinfo", func(ctx *gin.Context) {
		start := time.Now()
		cnf := config.Get()
		var stats syscall.Statfs_t
		err := syscall.Statfs(cnf.Global.DataDir, &stats)
		if err != nil {
			ctx.JSON(500, gin.H{
				"ok":    0,
				"errno": err.Error(),
			})
			return
		}

		total := stats.Blocks * uint64(stats.Bsize)
		used := total - stats.Bfree*uint64(stats.Bsize)

		ctx.JSON(200, gin.H{
			"ok": 1,
			"drive": gin.H{
				"used":  util.FSize(float64(used)),
				"total": util.FSize(float64(total)),
			},
			"respond_time": fmt.Sprintf("%dms", time.Since(start).Milliseconds()),
		})
	})
}
