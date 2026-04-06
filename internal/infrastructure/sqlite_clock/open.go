package sqlite_clock

import (
	"database/sql"
	"net/url"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func Open(path string) (*sql.DB, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	dsn := "file:" + url.PathEscape(absPath) + "?cache=shared"
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, err
	}

	return db, nil
}