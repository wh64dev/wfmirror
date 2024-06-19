package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/devproje/plog/level"
	"github.com/devproje/plog/log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/wh64dev/wfcloud/config"
	"github.com/wh64dev/wfcloud/routes"
	"github.com/wh64dev/wfcloud/service/auth"
	"github.com/wh64dev/wfcloud/util/database"
	"golang.org/x/term"
)

var (
	debug  bool
	server bool
)

func init() {
	flag.BoolVar(&debug, "D", false, "debug mode")
	flag.BoolVar(&server, "S", false, "run backend only")
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

	if _, err = os.ReadDir("./data"); err != nil {
		_ = os.Mkdir("data", 0775)
	}

	if _, err = os.ReadDir("./temp"); err != nil {
		_ = os.Mkdir("temp", 0775)
	}

	if _, err = os.ReadFile("./temp/config.json"); err != nil {
		cnf, _ := config.LoadDefault()
		_, _ = os.Create("temp/config.json")
		_ = os.WriteFile("./temp/config.json", cnf, 0755)
	}

	if _, err = os.ReadFile("./temp/service.db"); err != nil {
		_, _ = os.Create("temp/service.db")
	}
}

func main() {
	cnf := config.Get()
	app := gin.Default()
	database.Init()
	first()

	routes.New(app, server)

	fmt.Printf("Service bind port at http://localhost:%s\n", cnf.Port)
	fmt.Println("Mirror is now running. Press CTRL-C to exit.")

	err := app.Run(fmt.Sprintf(":%s", cnf.Port))
	if err != nil {
		log.Fatalln(err)
	}
}

func first() {
	accounts := auth.QueryAll()
	if len(accounts) > 0 {
		return
	}

	reader := bufio.NewReader(os.Stdin)

	// Read Username
	fmt.Print("Enter Username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}

	// Read Password
	fmt.Print("Enter Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println()

	fmt.Print("Enter Password one more time: ")
	pwCompare, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println()

	if string(bytePassword) != string(pwCompare) {
		log.Fatalln("typed password not compared")
	}

	acc := &auth.Account{
		Username: strings.TrimSpace(username),
		Password: string(bytePassword),
	}

	if _, err = acc.New(); err != nil {
		log.Fatalln(err)
	}
}
