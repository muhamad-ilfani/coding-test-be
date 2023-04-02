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

	tx, err := x.Postgresql.BeginTx(ctx, nil)
	if err == nil && tx != nil {
		defer func() { err = new(repository.SQLTransaction).EndTx(tx, err) }()
	}

	userPG := ur.NewRepository(tx)

	request := repository.GetAllTodosRequest{}

	response, httpcode, err := userPG.GetAllTodos(ctx, request)
	if err != nil {
		return res, httpcode, err
	}

	if len(res) <= 0 {
		return nil, httpcode, err
	}

	for _, val := range response {
		var deletedAt *string
		if str := val.DeletedAt.String(); str == "0001-01-01 00:00:00 +0000 UTC" {
			deletedAt = nil
		} else {
			deletedAt = &str
		}

		res = append(res, usecases.TodoList{
			ID:              val.ID,
			ActivityGroupID: val.ActivityGroupID,
			Title:           val.Title,
			IsActive:        val.IsActive,
			Priority:        val.Priority,
			CreatedAt:       val.CreatedAt.String(),
			UpdatedAt:       val.UpdatedAt.String(),
			DeletedAt:       deletedAt,
		})
	}

	return res, httpcode, err
}
