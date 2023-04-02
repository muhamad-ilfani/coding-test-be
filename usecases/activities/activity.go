package activity_case

import (
	"coding-test-be/usecases"
	"database/sql"
	"time"
)

func New(c Configuration, d Depencency) usecases.ActivityUseCase {
	return &usecase{c, d}
}

type Configuration struct {
	Timeout time.Duration
}

type Depencency struct {
	Postgresql *sql.DB
}

type usecase struct {
	Configuration
	Depencency
}
