package todo_case

import (
	"coding-test-be/repository"
	ur "coding-test-be/repository/todo_repository"
	"coding-test-be/usecases"
	"context"
	"errors"
	"net/http"
)

func (x *usecase) GetOneTodoByID(
	ctx context.Context, req usecases.GetOneTodoByIDRequest) (
	res usecases.GetOneTodoByIDResponse, httpcode int, err error,
) {
	ctx, cancel := context.WithTimeout(ctx, x.Configuration.Timeout)
	defer cancel()

	tx, err := x.Postgresql.BeginTx(ctx, nil)
	if err == nil && tx != nil {
		defer func() { err = new(repository.SQLTransaction).EndTx(tx, err) }()
	}

	activityPG := ur.NewRepository(tx)

	response, httpcode, err := activityPG.GetOneTodoByID(ctx, repository.GetOneTodoByIDRequest{ID: req.ID})
	if err != nil {
		return res, httpcode, err
	}
	if response.ID == 0 {
		return res, http.StatusNotFound, errors.New("id not found")
	}

	if str := response.DeletedAt.String(); str == "0001-01-01 00:00:00 +0000 UTC" {
		res.DeletedAt = nil
	} else {
		res.DeletedAt = &str
	}

	res = usecases.GetOneTodoByIDResponse{
		ID:              response.ID,
		ActivityGroupID: response.ActivityGroupID,
		Title:           response.Title,
		IsActive:        response.IsActive,
		Priority:        response.Priority,
		CreatedAt:       response.CreatedAt.String(),
		UpdatedAt:       response.UpdatedAt.String(),
	}

	return res, httpcode, err
}
