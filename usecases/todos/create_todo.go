package todo_case

import (
	"coding-test-be/repository"
	ur "coding-test-be/repository/todo_repository"
	"coding-test-be/usecases"
	"context"
)

func (x *usecase) CreateTodo(
	ctx context.Context, req usecases.CreateTodoRequest) (
	res usecases.CreateTodoResponse, httpcode int, err error,
) {
	ctx, cancel := context.WithTimeout(ctx, x.Configuration.Timeout)
	defer cancel()

	tx, err := x.Postgresql.BeginTxx(ctx, nil)
	if err == nil && tx != nil {
		defer func() { err = new(repository.SQLTransaction).EndTx(tx, err) }()
	}

	userPG := ur.NewRepository(tx)

	request := repository.CreateTodoRequest{
		Title:           req.Title,
		ActivityGroupID: req.ActivityGroupID,
		IsActive:        req.IsActive,
		Priority:        req.Priority,
	}

	response, httpcode, err := userPG.CreateTodo(ctx, request)
	if err != nil {
		return res, httpcode, err
	}

	res = usecases.CreateTodoResponse{
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
