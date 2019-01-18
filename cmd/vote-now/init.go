package main

import (
	"database/sql"
	"os"
)

func initDB() (*sql.DB, func()) {
	return NewPGConnection(os.Getenv("PG_DATASOURCE"))
}
