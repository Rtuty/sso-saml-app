package samlidp

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"modules/internal/tools"
)

type SqliteStore struct {
	*sql.DB
}

// NewSqliteStore создает новое хранилище на базе Sqlite
func NewSqliteStore(path string) (*SqliteStore, error) {
	isNeedCreate := !tools.FileExists(path)

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	if isNeedCreate {
		if _, err := db.Exec(`CREATE TABLE store (key TEXT UNIQUE, value TEXT);`); err != nil {
			return nil, err
		}
	}

	store := new(SqliteStore)
	store.DB = db

	return store, nil
}
