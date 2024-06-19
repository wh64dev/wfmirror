package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wh64dev/wfcloud/config"
)

type ConfigService struct{}

func (cs *ConfigService) LoadConfig(ctx *gin.Context) {
	cnf := config.Get()
	ctx.JSON(200, gin.H{
		"ok":     1,
		"status": 200,
		"data":   cnf.Global,
	})
}

func (cs *ConfigService) SetConfig(ctx *gin.Context) {
	t := ctx.PostForm("type")
	value := ctx.PostForm("value")

	switch t {
	case "dir":
		config.Set(t, value)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ok":     0,
			"status": http.StatusBadRequest,
			"errno":  "type name not matches",
		})
	}
}
