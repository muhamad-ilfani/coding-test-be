package app

import (
	"coding-test-be/delivery"
	"coding-test-be/usecases"
	activity_case "coding-test-be/usecases/activities"
	todo_case "coding-test-be/usecases/todos"
	"context"
	"time"
)

func (x *App) initService(ctx context.Context) {
	timeout := 55 * time.Second

	activitycase := activity_case.New(
		activity_case.Configuration{
			Timeout: timeout,
		},
		activity_case.Depencency{
			Postgresql: x.DB,
		},
	)

	todocase := todo_case.New(
		todo_case.Configuration{
			Timeout: timeout,
		},
		todo_case.Depencency{
			Postgresql: x.DB,
		},
	)

	delivery.NewEchoHandler(ctx, x.Echo, struct {
		usecases.ActivityUseCase
		usecases.TodoUseCase
	}{
		activitycase,
		todocase,
	})
}
