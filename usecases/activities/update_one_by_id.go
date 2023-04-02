package activity_case

import (
	"coding-test-be/repository"
	ur "coding-test-be/repository/activity_repository"
	"coding-test-be/usecases"
	"context"
	"errors"
	"net/http"
)

func (x *usecase) UpdateOneActivityByID(
	ctx context.Context, req usecases.UpdateOneActivityByIDRequest) (
	res usecases.UpdateOneActivityByIDResponse, httpcode int, err error,
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

	response, httpcode, err := activityPG.UpdateOneActivityByID(ctx, repository.UpdateOneActivityByIDRequest{
		ID:    req.ID,
		Title: req.Title,
	})
	if err != nil {
		return res, httpcode, err
	}

	if str := getData.DeletedAt.String(); str == "0001-01-01 00:00:00 +0000 UTC" {
		res.DeletedAt = nil
	} else {
		res.DeletedAt = &str
	}

	if req.Title != "" {
		getData.Title = req.Title
	}

	res = usecases.UpdateOneActivityByIDResponse{
		ID:        getData.ID,
		Title:     getData.Title,
		Email:     getData.Email,
		CreatedAt: getData.CreatedAt.String(),
		UpdatedAt: response.UpdatedAt.String(),
	}

	return res, httpcode, err
}
