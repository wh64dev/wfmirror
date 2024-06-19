package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wh64dev/wfcloud/service"
	"github.com/wh64dev/wfcloud/service/auth"
)

func CheckPriv(ctx *gin.Context) {
	path, _ := ctx.Params.Get("dirname")

	if strings.Contains(path, "/path/api/configuration") {
		_, validation := auth.Validate(ctx, true)
		if !validation {
			ctx.Abort()
			return
		}

		ctx.Next()
		return
	}

	priv := new(service.PrivDir)
	target, err := priv.Get(path)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return
		}

		ctx.JSON(500, gin.H{
			"ok":    0,
			"errno": err,
		})
		ctx.Abort()
		return
	}

	if strings.Contains(path, target.Path) {
		_, validation := auth.Validate(ctx, true)
		if !validation {
			ctx.Abort()
			return
		}
	}

	ctx.Next()
}
