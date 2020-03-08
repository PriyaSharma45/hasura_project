package service

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func PostgresConnection() (*sql.DB, error) {
	connStr := "postgres://postgres:mysecretpassword@localhost:5467/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return db, err
	}
	return db, nil
}
