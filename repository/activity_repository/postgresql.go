package activity_repo

import (
	"coding-test-be/repository"

	"github.com/jmoiron/sqlx"
)

type List []interface{}

type PostgreSQLConn struct {
	tc *sqlx.Tx
}

type Repository interface {
	repository.ActivityRepo
}

func NewRepository(tc *sqlx.Tx) Repository { return &PostgreSQLConn{tc} }
