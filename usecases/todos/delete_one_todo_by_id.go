package todo_case

import (
	"coding-test-be/repository"
	ur "coding-test-be/repository/todo_repository"
	"coding-test-be/usecases"
	"context"
)

func (x *usecase) DeleteOneTodoByID(
	ctx context.Context, req usecases.DeleteOneTodoByIDRequest) (
	res usecases.DeleteOneTodoByIDResponse, httpcode int, err error,
) {
	ctx, cancel := context.WithTimeout(ctx, x.Configuration.Timeout)
	defer cancel()

	tx, err := x.Postgresql.BeginTxx(ctx, nil)
	if err == nil && tx != nil {
		defer func() { err = new(repository.SQLTransaction).EndTx(tx, err) }()
	}

	activityPG := ur.NewRepository(tx)

	response, httpcode, err := activityPG.DeleteOneTodoByID(ctx, repository.DeleteOneTodoByIDRequest{ID: req.ID})
	if err != nil {
		return res, httpcode, err
	}

	res = usecases.DeleteOneTodoByIDResponse{
		ID: response.ID,
	}

	return res, httpcode, err
}
