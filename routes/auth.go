package routes

import (
	"github.com/devproje/plog/log"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/wh64dev/wfcloud/auth"
)

func CheckLogin(ctx *gin.Context) (string, int) {
	token, err := ctx.Request.Cookie("access-token")
	if err != nil {
		return "", 401
	}

	str := token.Value

	if str == "" {
		return "", 401
	}

	claims := &auth.Claims{}
	_, err = jwt.ParseWithClaims(str, claims, func(t *jwt.Token) (interface{}, error) {
		return auth.PrivKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", 401
		}

		log.Errorln(err)
		return "", 403
	}

	return claims.Id, 200
}
