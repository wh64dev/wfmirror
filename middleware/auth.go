package middleware

import (
	"fmt"
	"strings"

	"github.com/devproje/plog/log"
	"github.com/gin-gonic/gin"
	"github.com/wh64dev/wfcloud/util/database"
)

type privdir struct {
	Id   int
	Path string
}

func CheckPriv(ctx *gin.Context) {
	path := ctx.Request.URL.Path

	db := database.Open()
	stmt := `select * from privdir;`

	prep, err := db.Prepare(stmt)
	if err != nil {
		ctx.String(500, fmt.Sprintf("Internal Server Error\n%s\n", err))
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

	log.Infoln(path, data.Path)
	if strings.Contains(path, data.Path) {
		ctx.Abort()
		ctx.String(401, "Unauthorized")
		return
	}

	database.Close(db)
}
