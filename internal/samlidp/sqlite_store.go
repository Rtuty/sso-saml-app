package samlidp

import (
	"database/sql"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tenrok/saml/logger"
	"modules/internal/tools"
	"strings"
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

//Список возвращает все ключи, которые начинаются с `prefix`.
//Префикс удаляется из каждого возвращаемого значения. Например, если ключи ["aa", "ab", "cd"], то List("a") выдаст []строку {"a", "b"} */
func (s *SqliteStore) List(prefix string) ([]string, error) {
	rv := []string{}

	stmt, err := s.Prepare(`SELECT key FROM store  WHERE INSTR(key, :prefix) = 1`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(sql.Named("prefix", prefix))
	if err != nil && err != sql.ErrNoRows {
		return rv, err
	}
	defer rows.Close()

	for rows.Next() {
		var k string
		if err := rows.Scan(&k); err != nil {
			return rv, err
		}
		if strings.HasPrefix(k, prefix) {
			rv = append(rv, strings.TrimPrefix(k, prefix))
		}
	}

	return rv, nil
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

// Put todo
func (s *SqliteStore) Put(key string, value interface{}) error {
	tx, err := s.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT INTO store(key, value) VALUES(:key, :value) ON CONFLICT(key) DO UPDATE SET value = :value;`)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	v, _ := json.Marshal(value)
	if _, err := stmt.Exec(sql.Named("key", key), sql.Named("value", string(v))); err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}