package middleware

import (
	"strings"

	"github.com/devproje/plog/log"
	"github.com/gin-gonic/gin"
	"github.com/wh64dev/wfcloud/auth"
	"github.com/wh64dev/wfcloud/service"
)

func CheckPriv(ctx *gin.Context) {
	path, _ := ctx.Params.Get("dirname")

	priv := new(service.PrivDir)
	target, err := priv.Get(path)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return
		}

		ctx.String(500, "Database Error")
		log.Errorln(err)
		ctx.Abort()
		return
	}

	if strings.Contains(path, target.Path) {
		if !auth.Validate(ctx) {
			ctx.Abort()
			return
		}
	}

	ctx.Next()
}
