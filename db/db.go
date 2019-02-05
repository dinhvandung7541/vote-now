package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB
var CloseDB func()

// NewPGConnection create new postgres database connection
// it return postgres connection and cleanup postgres connection
func NewPGConnection(dataSource string) (*sql.DB, func()) {
	log.Println("env value: " + dataSource)
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db, func() {
		err := db.Close()
		if err != nil {
			log.Println("Failed to close DB by error", err)
		}
	}
}

func InitDBConnection() (*sql.DB, func()) {
	if DB == nil {
		DB, CloseDB = NewPGConnection(os.Getenv("PG_DATASOURCE"))
	}
	return DB, CloseDB
}
