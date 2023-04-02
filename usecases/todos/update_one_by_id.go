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
	if getData.ID == 0 || err != nil {
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

	if str := getData.DeletedAt.String(); str == "0001-01-01 00:00:00 +0000 UTC" {
		res.DeletedAt = nil
	} else {
		res.DeletedAt = &str
	}

	if req.Priority != "" {
		getData.Priority = req.Priority
	}
	if req.Title != "" {
		getData.Title = req.Title
	}
	if req.IsActive != getData.IsActive {
		getData.IsActive = req.IsActive
	}

	res = usecases.UpdateOneTodoByIDResponse{
		ID:              getData.ID,
		ActivityGroupID: getData.ActivityGroupID,
		Title:           getData.Title,
		IsActive:        getData.IsActive,
		Priority:        req.Priority,
		CreatedAt:       getData.CreatedAt.String(),
		UpdatedAt:       response.UpdatedAt.String(),
	}

	return res, httpcode, err
}
