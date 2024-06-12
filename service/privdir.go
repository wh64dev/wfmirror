package service

import (
	"github.com/devproje/plog/log"
	"github.com/wh64dev/wfcloud/util/database"
)

type PrivDir struct{}

type DirData struct {
	Id   int
	Path string
}

func (d *PrivDir) Get(dir string) (*DirData, error) {
	conn := database.Open()
	defer database.Close(conn)

	stmt := "select * from privdir where path like ?;"
	prep, err := conn.Prepare(stmt)
	if err != nil {
		return nil, err
	}

	var data DirData
	row := prep.QueryRow(dir)
	err = row.Scan(&data.Id, &data.Path)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (d *PrivDir) GetAll() ([]*DirData, error) {
	conn := database.Open()
	defer database.Close(conn)

	stmt := "select * from privdir;"
	prep, err := conn.Prepare(stmt)
	if err != nil {
		return nil, err
	}

	var dirs []*DirData
	rows, err := prep.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var data DirData
		err := rows.Scan(&data.Id, &data.Path)
		if err != nil {
			return nil, err
		}

		dirs = append(dirs, &data)
	}

	return dirs, nil
}

func (d *PrivDir) Add(dir string) error {
	conn := database.Open()
	defer database.Close(conn)

	stmt := "insert into privdir (path) values (?);"
	prep, err := conn.Prepare(stmt)
	if err != nil {
		return err
	}

	res, err := prep.Exec(dir)
	if err != nil {
		return err
	}

	aff, err := res.RowsAffected()
	if err != nil {
		return nil
	}

	log.Infof("Row affected count: %d\n", aff)
	return nil
}

func (d *PrivDir) Drop(id int) error {
	conn := database.Open()
	defer database.Close(conn)

	stmt := "delete from privdir where id = ?;"
	prep, err := conn.Prepare(stmt)
	if err != nil {
		return err
	}

	res, err := prep.Exec(id)
	if err != nil {
		return err
	}

	aff, err := res.RowsAffected()
	if err != nil {
		return nil
	}

	log.Infof("Row affected count: %d\n", aff)
	return nil
}
