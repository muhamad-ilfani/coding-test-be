package delivery

import (
	"coding-test-be/usecases"
	"context"

	"github.com/labstack/echo"
)

const (
	TokenIsRequired              = "Token must be provided"
	SuccessMsg                   = "Success"
	WelcomeMsg                   = "welcome"
	FailedToUnmarshall           = "Failed to Unmarshall"
	FailedToCreateActivity       = "Failed to Create Activity"
	FailedToGetDataActivities    = "Failed to Get Data Activitiy"
	FailedToUpdateDataActivities = "Failed to Update Data Activitiy"
	FailedToDeleteDataActivities = "Failed to Delete Data Activitiy"
	DeleteMsg                    = "Activity with ID %v Not Found"
)

type echoObject struct {
	*echo.Echo
	UseCase
}

type UseCase struct {
	usecases.ActivityUseCase
}

func NewEchoHandler(ctx context.Context, c *echo.Echo, uc UseCase) {
	obj := &echoObject{c, uc}
	obj.initRoute(ctx)

	obj.Logger.Fatal(obj.Start(":3030"))
}
