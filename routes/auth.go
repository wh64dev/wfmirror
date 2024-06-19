package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wh64dev/wfcloud/service/auth"
)

type AuthService struct{}

func (as *AuthService) Info(ctx *gin.Context) {
	claims, validation := auth.Validate(ctx, true)
	if !validation {
		return
	}

	ctx.JSON(200, gin.H{
		"ok":       1,
		"id":       claims.UserID,
		"username": claims.Username,
	})
}

func (as *AuthService) Register(ctx *gin.Context) {
	if !checkAuth(ctx) {
		ctx.JSON(401, gin.H{
			"ok":    0,
			"errno": "unauthorized access",
		})

		return
	}

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	acc := &auth.Account{
		Username: username,
		Password: password,
	}

	id, err := acc.New()
	if err != nil {
		ctx.JSON(500, gin.H{
			"ok":    0,
			"errno": err.Error(),
		})

		return
	}

	ctx.JSON(200, gin.H{
		"ok":     1,
		"action": "create",
		"id":     id,
	})
}

func (as *AuthService) Login(ctx *gin.Context) {
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
			"errno":  "username or password not matches!",
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

	ctx.JSON(200, gin.H{
		"ok":      1,
		"status":  200,
		"user_id": acc.Id,
		"token":   *token,
	})
}

func (as *AuthService) Accounts(ctx *gin.Context) {
	if !checkAuth(ctx) {
		ctx.JSON(401, gin.H{
			"ok":    0,
			"errno": "unauthorized access",
		})

		return
	}

	data := auth.QueryAll()
	ctx.JSON(200, gin.H{
		"ok":     1,
		"status": 200,
		"data":   data,
	})
}

func (as *AuthService) ChangePassword(ctx *gin.Context) {
	if !checkAuth(ctx) {
		ctx.JSON(401, gin.H{
			"ok":    0,
			"errno": "unauthorized access",
		})

		return
	}

	id := ctx.PostForm("id")
	password := ctx.PostForm("password")
	err := auth.ChangePassword(id, password)
	if err != nil {
		ctx.JSON(500, gin.H{
			"ok":    0,
			"errno": err.Error(),
		})
	}

	ctx.JSON(200, gin.H{
		"ok":     1,
		"action": "update",
		"id":     id,
	})
}

func (as *AuthService) DeleteAccount(ctx *gin.Context) {
	if !checkAuth(ctx) {
		ctx.JSON(401, gin.H{
			"ok":    0,
			"errno": "unauthorized access",
		})

		return
	}

	id := ctx.PostForm("id")
	err := auth.DeleteAccount(id)
	if err != nil {
		ctx.JSON(500, gin.H{
			"ok":    0,
			"errno": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"ok":     1,
		"action": "delete",
		"id":     id,
	})
}
