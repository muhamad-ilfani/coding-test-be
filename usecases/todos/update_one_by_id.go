package todo_case

import (
	"coding-test-be/repository"
	ur "coding-test-be/repository/todo_repository"
	"coding-test-be/usecases"
	"context"
)

func (x *usecase) UpdateOneTodoByID(
	ctx context.Context, req usecases.UpdateOneTodoByIDRequest) (
	res usecases.UpdateOneTodoByIDResponse, httpcode int, err error,
) {
	ctx, cancel := context.WithTimeout(ctx, x.Configuration.Timeout)
	defer cancel()

	tx, err := x.Postgresql.BeginTxx(ctx, nil)
	if err == nil && tx != nil {
		defer func() { err = new(repository.SQLTransaction).EndTx(tx, err) }()
	}

	activityPG := ur.NewRepository(tx)

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
