package todo_case

import (
	"coding-test-be/repository"
	ur "coding-test-be/repository/todo_repository"
	"coding-test-be/usecases"
	"context"
)

func (x *usecase) GetOneTodoByID(
	ctx context.Context, req usecases.GetOneTodoByIDRequest) (
	res usecases.GetOneTodoByIDResponse, httpcode int, err error,
) {
	ctx, cancel := context.WithTimeout(ctx, x.Configuration.Timeout)
	defer cancel()

	tx, err := x.Postgresql.BeginTxx(ctx, nil)
	if err == nil && tx != nil {
		defer func() { err = new(repository.SQLTransaction).EndTx(tx, err) }()
	}

	activityPG := ur.NewRepository(tx)

	response, httpcode, err := activityPG.GetOneTodoByID(ctx, repository.GetOneTodoByIDRequest{ID: req.ID})
	if err != nil {
		return res, httpcode, err
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
