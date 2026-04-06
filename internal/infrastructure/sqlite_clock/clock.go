package sqlite_clock

import (
	"database/sql"
	"time"
)

type SQLiteClock struct {
	db *sql.DB
}

func (c *SQLiteClock) getTimeFromCurrentTimestamp() time.Time {
	var ts string
	if err := c.db.QueryRow("SELECT CURRENT_TIMESTAMP;").Scan(&ts); err != nil {
		panic(err)
	}

	t, err := time.Parse(time.DateTime, ts)
	if err != nil {
		panic(err)
	}
	return t
}

func (c *SQLiteClock) NowUnix() int64 {
	t := c.getTimeFromCurrentTimestamp()
	return t.UTC().Unix()
}

func NewSQLiteClock(db *sql.DB) *SQLiteClock {
	return &SQLiteClock{db: db}
}
