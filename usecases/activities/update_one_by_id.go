package activity_case

import (
	"coding-test-be/repository"
	ur "coding-test-be/repository/user_repository"
	"coding-test-be/usecases"
	"context"
)

func (x *usecase) UpdateOneActivityByID(
	ctx context.Context, req usecases.UpdateOneActivityByIDRequest) (
	res usecases.UpdateOneActivityByIDResponse, httpcode int, err error,
) {
	ctx, cancel := context.WithTimeout(ctx, x.Configuration.Timeout)
	defer cancel()

	tx, err := x.Postgresql.BeginTxx(ctx, nil)
	if err == nil && tx != nil {
		defer func() { err = new(repository.SQLTransaction).EndTx(tx, err) }()
	}

	activityPG := ur.NewRepository(tx)

	response, httpcode, err := activityPG.UpdateOneActivityByID(ctx, repository.UpdateOneActivityByIDRequest{
		ID:    req.ID,
		Title: req.Title,
	})
	if err != nil {
		return res, httpcode, err
	}

	res = usecases.UpdateOneActivityByIDResponse{
		ID:        response.ID,
		Title:     response.Title,
		Email:     response.Email,
		CreatedAt: response.CreatedAt.String(),
		UpdatedAt: response.UpdatedAt.String(),
	}

	return res, httpcode, err
}
