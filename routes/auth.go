package routes

import (
	"strings"

	"github.com/devproje/plog/log"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Auth struct{}

func (a *Auth) Login(ctx *gin.Context) {
	username := ctx.Param("username")
	password := ctx.Param("password")

	id := uuid.New()
	log.Debugf("username: %s password: %s", username, strings.ReplaceAll(password, "", "*"))

	ctx.JSON(200, gin.H{
		"ok":     1,
		"status": 200,
		"id":     id.String(),
	})
}
