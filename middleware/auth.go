package middleware

import (
	"net/http"
	"strings"

	"github.com/devproje/plog/log"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/wh64dev/wfcloud/config"
	"github.com/wh64dev/wfcloud/util/database"
)

type privdir struct {
	Id   int
	Path string
}

func CheckPriv(ctx *gin.Context) {
	path := ctx.Request.URL.Path
	cnf := config.Get()

	db := database.Open()
	defer database.Close(db)

	stmt := `select * from privdir;`

	prep, err := db.Prepare(stmt)
	if err != nil {
		ctx.String(500, "Internal Server Error")
		log.Errorln(err)
		ctx.Abort()
		return
	}

	var data privdir
	row := prep.QueryRow()
	err = row.Scan(&data.Id, &data.Path)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return
		}

		ctx.String(500, "Database Error")
		log.Errorln(err)
		ctx.Abort()
		return
	}

	if strings.Contains(path, data.Path) {
		cookies := ctx.Request.Cookies()
		var tokenStr = ""
		for _, cookie := range cookies {
			if cookie.Name != "Authorization" {
				continue
			}

			tokenStr = cookie.Value
		}

		if tokenStr == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "request does not contain an access token"})
			ctx.Abort()
			return
		}

		claims := &jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return cnf.JWT.JWTToken, nil
		})

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			ctx.Abort()
			return
		}

		if !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			ctx.Abort()
			return
		}

		ctx.Next()
		return
	}
}
