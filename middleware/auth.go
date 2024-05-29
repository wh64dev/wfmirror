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

	// row, err := prep.Query(path)
	// if err != nil {
	// 	return
	// }

	var data privdir
	row := prep.QueryRow()
	err = row.Scan(&data.Id, &data.Path)
	if err != nil {
		ctx.String(500, "Database Error")
		ctx.Abort()
		return
	}

	log.Infoln(path, data.Path)
	if strings.Contains(path, data.Path) {
		ctx.Abort()
		ctx.String(401, "Unauthorized")
		return
	}

	// for row.Next() {
	// 	var data privdir
	// 	err = row.Scan(&data.Id, &data.Path)
	// 	if err != nil {
	// 		continue
	// 	}

	// 	ctx.String(200, path+" "+data.Path)
	// 	if strings.Contains(path, data.Path) {
	// 		// ctx.String(401, "Unauthorized")
	// 		ctx.Abort()
	// 		return
	// 	}

	// 	ctx.Abort()
	// }

	database.Close(db)
}
