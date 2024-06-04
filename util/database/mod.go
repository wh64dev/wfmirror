package database

import (
	"database/sql"

	"github.com/devproje/plog/log"
	_ "github.com/mattn/go-sqlite3"
)

func Init() {
	db := Open()
	var stmts []string

	stmts = append(stmts, `create table if not exists privdir(
		id   integer not null,
		path text,
		primary key(id)
	);`)

	stmts = append(stmts, `create table if not exists account(
		id  	 varchar not null,
		username varchar not null unique,
		password varchar not null,
		primary key(id)
	)`)

	for _, stmt := range stmts {
		prep, err := db.Prepare(stmt)
		if err != nil {
			log.Errorln(err)
			Close(db)
			return
		}

		_, err = prep.Exec()
		if err != nil {
			log.Errorln(err)
			Close(db)
			return
		}
	}

	Close(db)
}

func Open() *sql.DB {
	db, err := sql.Open("sqlite3", "./temp/service.db")
	if err != nil {
		log.Errorln(err)
		return nil
	}

	return db
}

func Close(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Errorln(err)
	}
}