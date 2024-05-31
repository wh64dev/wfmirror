package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/devproje/plog"
	"github.com/devproje/plog/level"
	"github.com/devproje/plog/log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/wh64dev/wfcloud/auth"
	"github.com/wh64dev/wfcloud/config"
	"github.com/wh64dev/wfcloud/routes"
	"github.com/wh64dev/wfcloud/util/database"
	"golang.org/x/term"
)

var (
	debug  bool
	single bool
)

func init() {
	flag.BoolVar(&debug, "D", false, "debug mode")
	flag.BoolVar(&single, "S", false, "run backend only")
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

	if _, err = os.ReadFile("./temp/service.db"); err != nil {
		_, _ = os.Create("temp/service.db")
	}
}

func task() {
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

	// Read password
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

	if err = acc.New(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	cnf := config.Get()
	app := gin.Default()
	database.Init()
	task()

	routes.New(app)

	if single {
		serve(app, cnf)
		return
	}

	go serve(app, cnf) // run backend

	var action = []string{"run", "dev"}
	if !debug {
		build(cnf)
		action = []string{"start"}
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
	process.Env = append(process.Env, fmt.Sprintf("FRONT_TITLE=%s", cnf.Frontend.Title))

	log.SetOutput(os.Stdout)
	process.Stdout = os.Stdout
	process.Stderr = os.Stderr

	if err := process.Run(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Frontend is now running. Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	process.Cancel()
}

func build(cnf *config.Config) {
	front := plog.New()
	front.Level = level.Info
	if debug {
		front.Level = level.Trace
	}

	fmt.Println("create next.js env file")
	os.Chdir("./frontend")
	file, err := os.Create(".env")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = file.Write([]byte(fmt.Sprintf("SERVER_PORT=%s\nFRONT_TITLE=%s\n", cnf.Port, cnf.Frontend.Title)))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("build next.js source...")
	process := exec.Command("pnpm", "build")
	if errors.Is(process.Err, exec.ErrDot) {
		process.Err = nil
	}

	front.SetOutput(process.Stdout)
	if err := process.Run(); err != nil {
		os.Remove(".env")
		os.Chdir("../")
		front.Fatalln(err)
	}

	os.Remove(".env")
	os.Chdir("../")
}

func serve(app *gin.Engine, cnf *config.Config) {
	fmt.Printf("Service bind port at %s\n", cnf.Port)
	err := app.Run(fmt.Sprintf(":%s", cnf.Port))
	if err != nil {
		log.Fatalln(err)
	}
}
