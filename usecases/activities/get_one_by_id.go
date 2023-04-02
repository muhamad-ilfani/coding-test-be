package activity_case

import (
	"coding-test-be/repository"
	ur "coding-test-be/repository/activity_repository"
	"coding-test-be/usecases"
	"context"
	"errors"
	"net/http"
)

func (x *usecase) GetOneActivityByID(
	ctx context.Context, req usecases.GetOneActivityByIDRequest) (
	res usecases.GetOneActivityByIDResponse, httpcode int, err error,
) {
	ctx, cancel := context.WithTimeout(ctx, x.Configuration.Timeout)
	defer cancel()

	tx, err := x.Postgresql.BeginTx(ctx, nil)
	if err == nil && tx != nil {
		defer func() { err = new(repository.SQLTransaction).EndTx(tx, err) }()
	}

	activityPG := ur.NewRepository(tx)

	response, httpcode, err := activityPG.GetOneActivityByID(ctx, repository.GetOneActivityByIDRequest{ID: req.ID})
	if err != nil {
		return res, httpcode, err
	}

	if response.ID == 0 {
		return res, http.StatusNotFound, errors.New("Not Found")
	}

	if str := response.DeletedAt.String(); str == "0001-01-01 00:00:00 +0000 UTC" {
		res.DeletedAt = nil
	} else {
		res.DeletedAt = &str
	}

	res = usecases.GetOneActivityByIDResponse{
		ID:        response.ID,
		Title:     response.Title,
		Email:     response.Email,
		CreatedAt: response.CreatedAt.String(),
		UpdatedAt: response.UpdatedAt.String(),
	}

	return res, httpcode, err
}
