package activity_case

import (
	"coding-test-be/repository"
	ur "coding-test-be/repository/user_repository"
	"coding-test-be/usecases"
	"context"
)

func (x *usecase) GetAllActivities(
	ctx context.Context, req usecases.GetAllActivitiesRequest) (
	res usecases.GetAllActivitiesResponse, httpcode int, err error,
) {
	ctx, cancel := context.WithTimeout(ctx, x.Configuration.Timeout)
	defer cancel()

	tx, err := x.Postgresql.BeginTxx(ctx, nil)
	if err == nil && tx != nil {
		defer func() { err = new(repository.SQLTransaction).EndTx(tx, err) }()
	}

	activityPG := ur.NewRepository(tx)

	response, httpcode, err := activityPG.GetAllActivities(ctx, repository.GetAllActivitiesRequest{})
	if err != nil {
		return res, httpcode, err
	}

	for _, val := range response {
		res = append(res, usecases.GetAllActivitiesData{
			ID:        val.ID,
			Title:     val.Title,
			Email:     val.Email,
			CreatedAt: val.CreatedAt.String(),
			UpdatedAt: val.UpdatedAt.String(),
		})
	}

	return res, httpcode, err
}
