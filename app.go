package main

import (
	"log"
	"text/template"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.SetFuncMap(template.FuncMap{})
	app.LoadHTMLGlob("static/*.html")

	// action := app.Group("/action")
	app.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})

	err := app.Run(":3000")
	if err != nil {
		log.Fatalln(err)
	}
}
