package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/wh64dev/wfcloud/config"
	"github.com/wh64dev/wfcloud/routes"
)

var debug bool

func init() {
	flag.BoolVar(&debug, "D", false, "debug mode")
	flag.Parse()

	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	if _, err = os.ReadDir("./data"); err != nil {
		_ = os.Mkdir("data", 0775)
	}

	if _, err = os.ReadDir("./temp"); err != nil {
		_ = os.Mkdir("temp", 0775)
	}
}

func main() {
	cnf := config.Get()
	app := gin.Default()
	routes.New(app)

	err := app.Run(fmt.Sprintf(":%s", cnf.Port))
	if err != nil {
		log.Fatalln(err)
	}
}
