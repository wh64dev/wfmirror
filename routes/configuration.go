package routes

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/wh64dev/wfcloud/config"
)

type ConfigService struct{}

func (cs *ConfigService) LoadConfig(ctx *gin.Context) {
	if !checkAuth(ctx) {
		ctx.JSON(401, gin.H{
			"ok":    0,
			"errno": "unauthorized access",
		})

		return
	}

	cnf := config.Get()
	ctx.JSON(200, gin.H{
		"ok":     1,
		"status": 200,
		"data":   cnf.Global,
	})
}

func (cs *ConfigService) SetConfig(ctx *gin.Context) {
	cnf := config.Get()
	if !checkAuth(ctx) {
		ctx.JSON(401, gin.H{
			"ok":    0,
			"errno": "unauthorized access",
		})

		return
	}

	t := ctx.PostForm("type")
	value := ctx.PostForm("value")

	switch t {
	case "dir":
		before := cnf.Global.DataDir
		config.Set(t, value)
		after := cnf.Global.DataDir
		err := os.Rename(before, after)
		if err != nil {
			ctx.JSON(500, gin.H{
				"ok":    0,
				"errno": err.Error(),
			})

			return
		}
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ok":     0,
			"status": http.StatusBadRequest,
			"errno":  "type name not matches",
		})
	}

	ctx.JSON(200, gin.H{
		"ok":   1,
		"type": t,
	})
}
