package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/wh64dev/wfcloud/config"
)

func CORS(ctx *gin.Context) {
	cnf := config.Get()
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", cnf.AllowOrigin)

	ctx.Next()
}
