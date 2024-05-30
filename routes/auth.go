package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Auth struct{}

func (a *Auth) Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	id := uuid.New()

	ctx.JSON(200, gin.H{
		"ok":       1,
		"status":   200,
		"id":       id.String(),
		"username": username,
		"password": password,
	})
}
