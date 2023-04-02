package activity_case

import (
	"coding-test-be/repository"
	ur "coding-test-be/repository/activity_repository"
	"coding-test-be/usecases"
	"context"
)

func (x *usecase) GetAllActivities(
	ctx context.Context, req usecases.GetAllActivitiesRequest) (
	res usecases.GetAllActivitiesResponse, httpcode int, err error,
) {
	ctx, cancel := context.WithTimeout(ctx, x.Configuration.Timeout)
	defer cancel()

	tx, err := x.Postgresql.BeginTx(ctx, nil)
	if err == nil && tx != nil {
		defer func() { err = new(repository.SQLTransaction).EndTx(tx, err) }()
	}

	activityPG := ur.NewRepository(tx)

	response, httpcode, err := activityPG.GetAllActivities(ctx, repository.GetAllActivitiesRequest{})
	if err != nil {
		return res, httpcode, err
	}

	for _, val := range response {
		var deletedAt *string
		if str := val.DeletedAt.String(); str == "0001-01-01 00:00:00 +0000 UTC" {
			deletedAt = nil
		} else {
			deletedAt = &str
		}

		res = append(res, usecases.GetAllActivitiesData{
			ID:        val.ID,
			Title:     val.Title,
			Email:     val.Email,
			CreatedAt: val.CreatedAt.String(),
			UpdatedAt: val.UpdatedAt.String(),
			DeletedAt: deletedAt,
		})
	}

	return res, httpcode, err
}
