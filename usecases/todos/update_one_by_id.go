package todo_case

import (
	"coding-test-be/repository"
	ur "coding-test-be/repository/todo_repository"
	"coding-test-be/usecases"
	"context"
	"errors"
	"net/http"
)

func (x *usecase) UpdateOneTodoByID(
	ctx context.Context, req usecases.UpdateOneTodoByIDRequest) (
	res usecases.UpdateOneTodoByIDResponse, httpcode int, err error,
) {
	ctx, cancel := context.WithTimeout(ctx, x.Configuration.Timeout)
	defer cancel()

	tx, err := x.Postgresql.BeginTx(ctx, nil)
	if err == nil && tx != nil {
		defer func() { err = new(repository.SQLTransaction).EndTx(tx, err) }()
	}

	activityPG := ur.NewRepository(tx)

	getData, httpcode, err := activityPG.GetOneTodoByID(ctx, repository.GetOneTodoByIDRequest{
		ID: req.ID,
	})
	if getData.Title == "" || err != nil {
		return res, http.StatusNotFound, errors.New("id not found")
	}

	response, httpcode, err := activityPG.UpdateOneTodoByID(ctx, repository.UpdateOneTodoByIDRequest{
		ID:       req.ID,
		Title:    req.Title,
		Priority: req.Priority,
		IsActive: req.IsActive,
	})
	if err != nil {
		return res, httpcode, err
	}

	res = usecases.UpdateOneTodoByIDResponse{
		ID:              getData.ID,
		ActivityGroupID: getData.ActivityGroupID,
		Title:           getData.Title,
		IsActive:        getData.IsActive,
		Priority:        getData.Priority,
		CreatedAt:       getData.CreatedAt.String(),
		UpdatedAt:       response.UpdatedAt.String(),
		DeletedAt:       getData.DeletedAt.String(),
	}

	return res, httpcode, err
}
