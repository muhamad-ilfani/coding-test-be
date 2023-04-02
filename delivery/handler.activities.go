package delivery

import (
	"coding-test-be/usecases"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllActivities(ctx context.Context, uc usecases.ActivityUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx = c.Request().WithContext(ctx).Context()

		form := usecases.GetAllActivitiesRequest{}

		res, httpcode, err := uc.GetAllActivities(ctx, form)
		if err != nil {
			return c.JSON(httpcode, map[string]interface{}{
				"message": FailedToGetDataActivities,
				"error":   err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  SuccessMsg,
			"message": SuccessMsg,
			"data":    res,
		})
	}
}

func CreateActivity(ctx context.Context, uc usecases.ActivityUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx = c.Request().WithContext(ctx).Context()

		form := usecases.CreateActivityRequest{}

		err := c.Bind(&form)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": FailedToUnmarshall,
				"error":   err.Error(),
			})
		}

		res, httpcode, err := uc.CreateActivity(ctx, form)
		if err != nil {
			return c.JSON(httpcode, map[string]interface{}{
				"status":  "Bad Request",
				"message": err.Error(),
				"data":    nil,
			})
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"status":  SuccessMsg,
			"message": SuccessMsg,
			"data":    res,
		})
	}
}

func GetOneByID(ctx context.Context, uc usecases.ActivityUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx = c.Request().WithContext(ctx).Context()

		GetId, _ := strconv.Atoi(c.Param("id"))

		form := usecases.GetOneActivityByIDRequest{
			ID: int64(GetId),
		}

		res, httpcode, err := uc.GetOneActivityByID(ctx, form)
		if err != nil {
			return c.JSON(httpcode, map[string]interface{}{
				"status":  "Not Found",
				"message": fmt.Sprintf(DeleteMsg, GetId),
				"data":    nil,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  SuccessMsg,
			"message": SuccessMsg,
			"data":    res,
		})
	}
}

func UpdateOneActivityByID(ctx context.Context, uc usecases.ActivityUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx = c.Request().WithContext(ctx).Context()
		form := usecases.UpdateOneActivityByIDRequest{}

		GetId, _ := strconv.Atoi(c.Param("id"))

		err := c.Bind(&form)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": FailedToUnmarshall,
				"error":   err.Error(),
			})
		}

		form.ID = int64(GetId)

		res, httpcode, err := uc.UpdateOneActivityByID(ctx, form)
		if err != nil {
			return c.JSON(httpcode, map[string]interface{}{
				"status":  "Not Found",
				"message": fmt.Sprintf(DeleteMsg, GetId),
				"data":    nil,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  SuccessMsg,
			"message": SuccessMsg,
			"data":    res,
		})
	}
}

func DeleteOneActivityByID(ctx context.Context, uc usecases.ActivityUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx = c.Request().WithContext(ctx).Context()

		GetId, _ := strconv.Atoi(c.Param("id"))

		form := usecases.DeleteOneActivityByIDRequest{ID: int64(GetId)}

		_, httpcode, err := uc.DeleteOneActivityByID(ctx, form)
		if err != nil {
			return c.JSON(httpcode, map[string]interface{}{
				"status":  "Not Found",
				"message": fmt.Sprintf(DeleteMsg, GetId),
				"data":    nil,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "Success",
			"message": fmt.Sprintf(DeleteMsg, GetId),
		})
	}
}
