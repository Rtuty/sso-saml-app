package samlidp

import (
	"database/sql"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tenrok/saml/logger"
	"modules/internal/tools"
)

type SqliteStore struct {
	*sql.DB
	logger logger.Interface
}

// NewSqliteStore создает новое хранилище на базе Sqlite
func NewSqliteStore(path string) (*SqliteStore, error) {
	isNeedCreate := !tools.FileExists(path)

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		logger.DefaultLogger.Printf("SqLite doesn't opened. ERROR: %s", err)
		return nil, err
	}

	if isNeedCreate {
		if _, err := db.Exec(`
			CREATE TABLE store 
			(
			    key TEXT UNIQUE,
			    value TEXT
			);`); err != nil {
			return nil, err
		}
	}

	store := new(SqliteStore)
	store.DB = db

	return store, nil
}

// Get достает ключи из БД и возвращает value
func (s *SqliteStore) Get(key string, value interface{}) error {
	stmt, err := s.Prepare(`SELECT value FROM store WHERE key = :key`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var v string
	if err := stmt.QueryRow(sql.Named("key", key)).Scan(&v); err != nil {
		if err == sql.ErrNoRows {
			s.logger.Printf("Rows not found. ERROR: %s", err)
			return err
		}
	}

	return json.Unmarshal([]byte(v), &value)
}