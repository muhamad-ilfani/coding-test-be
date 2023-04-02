package activity_repo

import (
	"coding-test-be/repository"
	"database/sql"
)

type List []interface{}

type PostgreSQLConn struct {
	tc *sql.Tx
}

type Repository interface {
	repository.ActivityRepo
}

func NewRepository(tc *sql.Tx) Repository { return &PostgreSQLConn{tc} }
