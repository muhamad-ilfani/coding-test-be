package activity_case

import (
	"coding-test-be/repository"
	ur "coding-test-be/repository/activity_repository"
	"coding-test-be/usecases"
	"context"
	"errors"
	"net/http"
)

func (x *usecase) DeleteOneActivityByID(
	ctx context.Context, req usecases.DeleteOneActivityByIDRequest) (
	res usecases.DeleteOneActivityByIDResponse, httpcode int, err error,
) {
	ctx, cancel := context.WithTimeout(ctx, x.Configuration.Timeout)
	defer cancel()

	tx, err := x.Postgresql.BeginTx(ctx, nil)
	if err == nil && tx != nil {
		defer func() { err = new(repository.SQLTransaction).EndTx(tx, err) }()
	}

	activityPG := ur.NewRepository(tx)

	getData, httpcode, err := activityPG.GetOneActivityByID(ctx, repository.GetOneActivityByIDRequest{
		ID: req.ID,
	})
	if getData.ID == 0 {
		return res, http.StatusNotFound, errors.New("id not found")
	}

	response, httpcode, err := activityPG.DeleteOneActivityByID(ctx, repository.DeleteOneActivityByIDRequest{ID: req.ID})
	if err != nil {
		return res, httpcode, err
	}

	res = usecases.DeleteOneActivityByIDResponse{
		ID: response.ID,
	}

	return res, httpcode, err
}
