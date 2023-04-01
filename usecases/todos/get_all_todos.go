package todo_case

import (
	"coding-test-be/repository"
	ur "coding-test-be/repository/todo_repository"
	"coding-test-be/usecases"
	"context"
)

func (x *usecase) GetAllTodos(
	ctx context.Context, req usecases.GetAllTodosRequest) (
	res usecases.GetAllTodosResponse, httpcode int, err error,
) {
	ctx, cancel := context.WithTimeout(ctx, x.Configuration.Timeout)
	defer cancel()

	tx, err := x.Postgresql.BeginTxx(ctx, nil)
	if err == nil && tx != nil {
		defer func() { err = new(repository.SQLTransaction).EndTx(tx, err) }()
	}

	userPG := ur.NewRepository(tx)

	request := repository.GetAllTodosRequest{}

	response, httpcode, err := userPG.GetAllTodos(ctx, request)
	if err != nil {
		return res, httpcode, err
	}

	for _, val := range response {
		res = append(res, usecases.TodoList{
			ID:              val.ID,
			ActivityGroupID: val.ActivityGroupID,
			Title:           val.Title,
			IsActive:        val.IsActive,
			Priority:        val.Priority,
			CreatedAt:       val.CreatedAt.String(),
			UpdatedAt:       val.UpdatedAt.String(),
		})
	}

	return res, httpcode, err
}
