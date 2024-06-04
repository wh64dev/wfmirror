package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wh64dev/wfcloud/auth"
)

type AuthService struct{}

func (as *AuthService) Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if username != "root" || password != "asdf1234!" { // TODO: example
		ctx.JSON(401, gin.H{
			"ok":     0,
			"status": 401,
			"errno":  "username or password not matches!",
		})

		return
	}

	acc := auth.AccountData{
		Id:       "wheel",
		Username: "root",
	}

	token, err := acc.GenToken()
	if err != nil {
		ctx.JSON(500, gin.H{
			"ok":     0,
			"status": 500,
			"errno":  err.Error(),
		})

		return
	}

	ctx.JSON(200, gin.H{
		"ok":     1,
		"status": 200,
		"token":  *token,
	})
}
