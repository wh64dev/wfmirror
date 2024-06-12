package routes

import (
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/wh64dev/wfcloud/config"
	"github.com/wh64dev/wfcloud/middleware"
)

func New(app *gin.Engine, server bool) {
	cnf := config.Get()
	dirWorker := new(DirWorker)
	as := new(AuthService)

	app.Use(middleware.CORS)
	app.Use(middleware.CheckPriv)

	if !server {
		app.Static("/static", "./static")
		app.LoadHTMLGlob("./pages/*.html")

		app.GET("/", func(ctx *gin.Context) {
			base := filepath.Join(uploadBaseDir, "")
			entries, err := worker(base, "")
			if err != nil {
				return
			}

			var temp = ""
			for _, entry := range entries {
				temp += fmt.Sprintf("<a href=\"%s\">%s</a><br/>", entry.URL, entry.Name)
			}

			ctx.HTML(200, "page.html", gin.H{
				"name": cnf.Service.Name,
				"dir":  template.HTML(temp),
			})
		})

		app.GET("/login", func(ctx *gin.Context) {
			ctx.HTML(200, "login.html", gin.H{
				"name": cnf.Service.Name,
			})
		})
	}

	api := app.Group("/api")
	{
		path := api.Group("/path")
		{
			path.GET("/*dirname", dirWorker.List)
			path.PUT("/upload/*dirname", dirWorker.UploadFile)
			path.POST("/secret/*dirname", dirWorker.AddSecret)
			path.DELETE("/secret/:id", dirWorker.DropSecret)
		}

		auth := api.Group("/auth")
		{
			auth.GET("/", as.Info)
			auth.POST("/login", as.Login)
		}
	}
}
