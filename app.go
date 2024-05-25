package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

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
}

func main() {
	cnf := config.Get()
	app := gin.Default()
	routes.New(app)

	go frontend(cnf)
	err := app.Run(fmt.Sprintf(":%s", cnf.Port))
	if err != nil {
		log.Fatalln(err)
	}
}

func frontend(cnf *config.Config) {
	var action = []string{"start"}
	if debug {
		action = []string{"run", "dev"}
	}

	command := []string{"-C", "./frontend"}

	command = append(command, action...)
	command = append(command, "--hostname")
	command = append(command, cnf.Frontend.Host)
	command = append(command, "--port")
	command = append(command, cnf.Frontend.Port)

	process := exec.Command("pnpm", command...)
	if errors.Is(process.Err, exec.ErrDot) {
		process.Err = nil
	}

	process.Env = os.Environ()
	process.Env = append(process.Env, fmt.Sprintf("SERVER_PORT=%s", cnf.Port))

	process.Stdout = os.Stdout
	process.Stderr = os.Stderr

	if err := process.Run(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Frontend is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	log.Println("close front server")
	process.Cancel()
}
