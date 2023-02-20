package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"uploadService/cmd/server/config"
	"uploadService/libs/errors"
)

type Database interface {
	CreateImageInfo(name string) error
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

func (db DatabaseConn) CreateImageInfo(name string) error {
	return nil
}
