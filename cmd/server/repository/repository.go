package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"uploadService/cmd/server/config"
	"uploadService/libs/errors"
)

type Database interface {
}

type DatabaseConn struct {
	conn *sql.DB
}

func NewDatabase() (*DatabaseConn, error) {
	conf := config.UploadServerConfig()

	connStr := conf.PgConnString

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		errors.HandleFatalError(err, "failed to connect to postgres")
	}

	err = db.Ping()

	if err != nil {
		errors.HandleFatalError(err, "failed to connect to postgres")
	}

	return &DatabaseConn{conn: db}, err
}
