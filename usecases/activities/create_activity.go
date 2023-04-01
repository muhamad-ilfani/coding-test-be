package activity_case

import (
	"coding-test-be/repository"
	"coding-test-be/usecases"
	"context"

	ur "coding-test-be/repository/activity_repository"
)

func (x *usecase) CreateActivity(
	ctx context.Context, req usecases.CreateActivityRequest) (
	res usecases.CreateActivityResponse, httpcode int, err error,
) {
	ctx, cancel := context.WithTimeout(ctx, x.Configuration.Timeout)
	defer cancel()

	tx, err := x.Postgresql.BeginTxx(ctx, nil)
	if err == nil && tx != nil {
		defer func() { err = new(repository.SQLTransaction).EndTx(tx, err) }()
	}

	userPG := ur.NewRepository(tx)

	request := repository.CreateActivityRequest{
		Title: req.Title,
		Email: req.Email,
	}

	response, httpcode, err := userPG.CreateActivity(ctx, request)
	if err != nil {
		return res, httpcode, err
	}

	res = usecases.CreateActivityResponse{
		ID:        response.ID,
		Title:     response.Title,
		Email:     response.Email,
		CreatedAt: response.CreatedAt.String(),
		UpdatedAt: response.UpdatedAt.String(),
	}

	return res, httpcode, err
}
