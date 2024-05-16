package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/wh64dev/wfcloud/routes"
)

func main() {
	if _, err := os.ReadDir("./data"); err != nil {
		_ = os.Mkdir("data", 0775)
	}

	app := gin.Default()
	routes.New(app)

	err := app.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
