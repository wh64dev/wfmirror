package auth

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/devproje/plog/log"
	"github.com/wh64dev/wfcloud/util"
	"golang.org/x/term"
)

const (
	accFilename = "wfconf/account.json"
)

type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func register() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print("Enter Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return err
	}

	fmt.Println()

	acc := &Account{
		Username: strings.TrimSpace(username),
		Password: string(bytePassword),
	}

	err = acc.new()
	if err != nil {
		return err
	}

	return nil
}

func (acc *Account) new() error {
	acc.Password = HashPassword(acc.Password)

	data, err := json.Marshal(acc)
	if err != nil {
		return err
	}

	err = os.WriteFile(accFilename, data, 0655)
	if err != nil {
		return err
	}

	return nil
}

func get() (*Account, error) {
	data, err := util.ParseJSON[Account](accFilename)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func InitAuth() {
	_, err := os.Stat(accFilename)
	if err != nil {
		err = register()
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func (acc *Account) Login() bool {
	info, err := get()
	if err != nil {
		return false
	}

	if info.Username != acc.Username {
		return false
	}

	return CheckPasswordHash(acc.Password, info.Password)
}
