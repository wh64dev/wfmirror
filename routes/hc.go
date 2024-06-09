package routes

import "github.com/gin-gonic/gin"

func HC(ctx *gin.Context) {
	ctx.Status(200)
}
