package database

import (
	"database/sql"
	"forum/config"

	_ "github.com/mattn/go-sqlite3"
)

const (
	userTable = `CREATE TABLE IF NOT EXISTS user (
		id  INTEGER PRIMARY KEY AUTOINCREMENT,
		username text,
		password text,
		confirm_password text,
		email text,
		post int,
		creation_time text,
		session_token text,
		expiration_token DATETIME DEFAULT NULL);
`
	postTable = `CREATE TABLE IF NOT EXISTS posts(
		id SERIAL PRIMARY KEY,
		context text,
		author varchar(255));
`
)

func InitDB(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open(cfg.DbDriver, cfg.DbNameAndPath)
	if err != nil {
		return nil, err
	}
	// checks connection to db
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func CreateTables(db *sql.DB) error {
	tables := []string{userTable, postTable}
	for _, table := range tables {
		_, err := db.Exec(table)
		if err != nil {
			return err
		}
	}
	return nil
}
