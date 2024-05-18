package routes

import (
	"fmt"
	"html/template"
	"time"

	"github.com/devproje/plog/log"
	"github.com/gin-gonic/gin"
	"github.com/wh64dev/wfcloud/auth"
)

func New(app *gin.Engine) {
	app.SetFuncMap(template.FuncMap{})
	app.LoadHTMLGlob("pages/*.html")
	app.Static("/static", "./static")

	action := app.Group("/action")
	{
		action.POST("/upload", func(ctx *gin.Context) {
			form, err := ctx.MultipartForm()
			if err != nil {
				return
			}

			var path string
			raw := form.Value["path"]
			files := form.File["upload[]"]
			if raw[0] == "/" {
				path = ""
			}

			for _, file := range files {
				log.Printf("received upload file: /%s/%s\n", path, file.Filename)
				ctx.SaveUploadedFile(file, fmt.Sprintf("./data%s/%s", path, file.Filename))
			}

			ctx.JSON(200, gin.H{
				"ok":      1,
				"status":  200,
				"count":   len(files),
				"workdir": fmt.Sprintf("/%s", path),
			})
		})

		action.POST("/login", func(ctx *gin.Context) {
			var username, password string
			var chk bool

			username, chk = ctx.GetPostForm("password")
			if !chk {
				return
			}

			password, chk = ctx.GetPostForm("password")
			if !chk {
				return
			}

			acc := &auth.Account{
				Username: username,
				Password: password,
			}

			chk = acc.Login()
			if !chk {
				return
			}

			token, err := acc.GenToken()
			if err != nil {
				return
			}

			ctx.Header("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-cache=0, max-age=0")
			ctx.Header("Last-Modified", time.Now().String())
			ctx.Header("Pragma", "no-cache")
			ctx.Header("Expires", "-1")

			ctx.SetCookie("access-token", token, 1800, "", "", false, false)

			ctx.Redirect(301, "/")
		})
	}

	app.GET("/", func(ctx *gin.Context) {
		DirWorker(ctx, "/")
	})

	app.GET("/:path/*child", func(ctx *gin.Context) {
		origin := ctx.Param("path")
		child := ctx.Param("child")
		var path = origin
		if child != "/" {
			path = fmt.Sprintf("%s%s", origin, child)
		}

		if auth.CheckAuth(origin) {
			ctx.HTML(401, "auth.html", gin.H{
				"path": fmt.Sprintf("/%s", origin),
			})
			return
		}

		DirWorker(ctx, path)
	})
}
