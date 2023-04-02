package activity_case

import (
	"coding-test-be/repository"
	"coding-test-be/usecases"
	"context"
	"errors"
	"net/http"

	ur "coding-test-be/repository/activity_repository"
)

func (x *usecase) CreateActivity(
	ctx context.Context, req usecases.CreateActivityRequest) (
	res usecases.CreateActivityResponse, httpcode int, err error,
) {
	ctx, cancel := context.WithTimeout(ctx, x.Configuration.Timeout)
	defer cancel()

	if req.Title == "" {
		return res, http.StatusBadRequest, errors.New("title must be provide")
	}

	tx, err := x.Postgresql.BeginTx(ctx, nil)
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

	getLatesID, httpcode, err := userPG.GetLatesActivityID(ctx, repository.GetLatesActivityIDRequest{})
	if err != nil {
		return res, httpcode, err
	}

	res = usecases.CreateActivityResponse{
		ID:        getLatesID.ID,
		Title:     response.Title,
		Email:     response.Email,
		CreatedAt: response.CreatedAt.String(),
		UpdatedAt: response.UpdatedAt.String(),
	}

	return res, httpcode, err
}
