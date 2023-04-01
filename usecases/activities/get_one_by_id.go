package activity_case

import (
	"coding-test-be/repository"
	ur "coding-test-be/repository/user_repository"
	"coding-test-be/usecases"
	"context"
)

func (x *usecase) GetOneActivityByID(
	ctx context.Context, req usecases.GetOneActivityByIDRequest) (
	res usecases.GetOneActivityByIDResponse, httpcode int, err error,
) {
	ctx, cancel := context.WithTimeout(ctx, x.Configuration.Timeout)
	defer cancel()

	tx, err := x.Postgresql.BeginTxx(ctx, nil)
	if err == nil && tx != nil {
		defer func() { err = new(repository.SQLTransaction).EndTx(tx, err) }()
	}

	activityPG := ur.NewRepository(tx)

	response, httpcode, err := activityPG.GetOneActivityByID(ctx, repository.GetOneActivityByIDRequest{ID: req.ID})
	if err != nil {
		return res, httpcode, err
	}

	res = usecases.GetOneActivityByIDResponse{
		ID:        response.ID,
		Title:     response.Title,
		Email:     response.Email,
		CreatedAt: response.CreatedAt.String(),
		UpdatedAt: response.UpdatedAt.String(),
	}

	return res, httpcode, err
}
