package activity_case

import (
	"coding-test-be/usecases"
	"time"

	"github.com/jmoiron/sqlx"
)

func New(c Configuration, d Depencency) usecases.ActivityUseCase {
	return &usecase{c, d}
}

type Configuration struct {
	Timeout time.Duration
}

type Depencency struct {
	Postgresql *sqlx.DB
}

type usecase struct {
	Configuration
	Depencency
}
