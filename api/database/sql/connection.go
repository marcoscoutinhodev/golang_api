package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@/")

	if err != nil {
		return nil, err
	}

	return db, nil
}
