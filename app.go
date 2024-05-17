package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/wh64dev/wfcloud/auth"
	"github.com/wh64dev/wfcloud/routes"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	if _, err = os.ReadDir("./wfconf"); err != nil {
		_ = os.Mkdir("wfconf", 0755)
	}

	if _, err = os.ReadDir("./data"); err != nil {
		_ = os.Mkdir("data", 0775)
	}

	auth.InitAuth()
	auth.InitFile()
}

func main() {
	app := gin.Default()
	routes.New(app)

	err := app.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatalln(err)
	}
}
