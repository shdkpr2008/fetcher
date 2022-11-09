package sqlite

import (
	"database/sql"
	"fetcher/internal/config"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const tableSchema = `CREATE TABLE IF NOT EXISTS metadata (site text PRIMARY KEY, 
														  num_links integer, images integer, last_fetch text);`

type SQLite struct {
	*sql.DB
}

func NewSQLite(config config.Config) *SQLite {
	sqlite := &SQLite{}
	db, err := sql.Open("sqlite3", config.DatabaseFile())
	if err != nil {
		log.Fatal(err)
	}
	sqlite.DB = db

	_, err = db.Prepare(tableSchema)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(tableSchema)
	if err != nil {
		log.Fatal(err)
	}

	return sqlite
}
