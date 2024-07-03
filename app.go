package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/devproje/plog/level"
	"github.com/devproje/plog/log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/wh64dev/wfcloud/config"
	"github.com/wh64dev/wfcloud/routes"
)

var debug bool

func init() {
	flag.BoolVar(&debug, "D", false, "debug mode")
	flag.Parse()

	log.SetLevel(level.Info)

	if !debug {
		log.SetLevel(level.Trace)
		gin.SetMode(gin.ReleaseMode)
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	cnf := config.Get()
	if _, err = os.ReadDir(cnf.Global.DataDir); err != nil {
		_ = os.Mkdir(cnf.Global.DataDir, 0775)
	}
}

func main() {
	cnf := config.Get()
	app := gin.Default()

	routes.New(app)

	port, err := strconv.ParseInt(cnf.Port, 10, 32)
	if err != nil {
		log.Fatalln(err)
	}

	if !debug {
		fmt.Printf("Service bind port at http://localhost:%s\n", cnf.Port)
		fmt.Println("Mirror is now running. Press CTRL-C to exit.")
	}

	err = Run(app, fmt.Sprintf(":%d", port))
	if err != nil {
		if !debug {
			log.Fatalln(err)
		}

		for err != nil {
			port += 1
			err = Run(app, fmt.Sprintf(":%d", port))
		}
	}
}

func Run(app *gin.Engine, port string) error {
	return app.Run(port)
}
