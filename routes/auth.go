package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wh64dev/wfcloud/auth"
)

type Auth struct{}

func (a *Auth) Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	authData := &auth.AuthForm{
		Username: username,
		Password: password,
	}

	acc, err := authData.Login()
	if err != nil {
		ctx.JSON(401, gin.H{
			"ok":     0,
			"status": 401,
			"errno":  err.Error(),
		})

		return
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

	ctx.Request.AddCookie(&http.Cookie{
		Name:  "Authorization",
		Value: token,
	})
	ctx.JSON(200, gin.H{
		"ok":     1,
		"status": 200,
	})
}
