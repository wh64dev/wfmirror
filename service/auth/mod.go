package auth

import (
	"errors"

	"github.com/devproje/plog/log"
	"github.com/google/uuid"
	"github.com/wh64dev/wfcloud/util/database"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	Id       string
	Username string
	Password string
}

type AuthForm struct {
	Username string
	Password string
}

func Hash(password string) string {
	pwd, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(pwd)
}

func CheckHash(pwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}

func (acc *Account) New() error {
	newID := uuid.New()
	hashedPw := Hash(acc.Password)

	db := database.Open()
	defer database.Close(db)

	stmt := "insert into account values (?, ?, ?);"
	prep, err := db.Prepare(stmt)
	if err != nil {
		return err
	}

	res, err := prep.Exec(newID.String(), acc.Username, hashedPw)
	if err != nil {
		return err
	}

	id, _ := res.LastInsertId()
	log.Infof("row inserted id: %d\n", id)

	return nil
}

func (af *AuthForm) Login() (*Account, error) {
	db := database.Open()
	defer database.Close(db)

	stmt := "select id, username, password from account where username = ?;"
	prep, err := db.Prepare(stmt)
	if err != nil {
		return nil, err
	}

	var data Account

	err = prep.QueryRow(af.Username).Scan(&data.Id, &data.Username, &data.Password)
	if err != nil {
		return nil, err
	}

	if !CheckHash(af.Password, data.Password) {
		return nil, errors.New("password not matches")
	}

	return &data, nil
}

func QueryAll() []*Account {
	db := database.Open()
	defer database.Close(db)

	stmt := "select * from account;"

	prep, err := db.Prepare(stmt)
	if err != nil {
		return nil
	}

	res, err := prep.Query()
	if err != nil {
		return nil
	}

	var accounts []*Account
	for res.Next() {
		var data Account
		err = res.Scan(&data.Id, &data.Username, &data.Password)
		if err != nil {
			continue
		}

		accounts = append(accounts, &data)
	}

	return accounts
}
