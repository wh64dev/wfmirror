package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wh64dev/wfcloud/middleware"
)

func New(app *gin.Engine) {
	worker := new(DirWorker)
	as := new(AuthService)

	app.Use(middleware.CORS)
	app.Use(middleware.CheckPriv)

	api := app.Group("/api")
	{
		api.GET("/path/*dirname", worker.List)
		api.POST("/upload/*dirname", worker.UploadFile)

		auth := api.Group("/auth")
		{
			auth.POST("/login", as.Login)
		}
	}
}
