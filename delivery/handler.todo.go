package delivery

import (
	"context"
	"net/http"
	"strconv"

	"coding-test-be/usecases"

	"github.com/labstack/echo"
)

func CreateTodo(ctx context.Context, uc usecases.TodoUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx = c.Request().WithContext(ctx).Context()

		form := usecases.CreateTodoRequest{}

		err := c.Bind(&form)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": FailedToUnmarshall,
				"error":   err.Error(),
			})
		}

		res, httpcode, err := uc.CreateTodo(ctx, form)
		if err != nil {
			return c.JSON(httpcode, map[string]interface{}{
				"message": FailedToCreateActivity,
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

func GetAllTodos(ctx context.Context, uc usecases.TodoUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx = c.Request().WithContext(ctx).Context()

		form := usecases.GetAllTodosRequest{}

		res, httpcode, err := uc.GetAllTodos(ctx, form)
		if err != nil {
			return c.JSON(httpcode, map[string]interface{}{
				"message": FailedToCreateActivity,
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

func GetOneTodoByID(ctx context.Context, uc usecases.TodoUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx = c.Request().WithContext(ctx).Context()

		GetId, _ := strconv.Atoi(c.Param("id"))

		form := usecases.GetOneTodoByIDRequest{ID: int64(GetId)}

		res, httpcode, err := uc.GetOneTodoByID(ctx, form)
		if err != nil {
			return c.JSON(httpcode, map[string]interface{}{
				"message": FailedToCreateActivity,
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

func UpdateOneTodoByID(ctx context.Context, uc usecases.TodoUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx = c.Request().WithContext(ctx).Context()
		form := usecases.UpdateOneTodoByIDRequest{}

		GetId, _ := strconv.Atoi(c.Param("id"))

		err := c.Bind(&form)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": FailedToUnmarshall,
				"error":   err.Error(),
			})
		}

		form.ID = int64(GetId)

		res, httpcode, err := uc.UpdateOneTodoByID(ctx, form)
		if err != nil {
			return c.JSON(httpcode, map[string]interface{}{
				"message": FailedToUpdateDataActivities,
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

func DeleteOneTodoByID(ctx context.Context, uc usecases.TodoUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx = c.Request().WithContext(ctx).Context()

		GetId, _ := strconv.Atoi(c.Param("id"))

		form := usecases.DeleteOneTodoByIDRequest{ID: int64(GetId)}

		_, httpcode, err := uc.DeleteOneTodoByID(ctx, form)
		if err != nil {
			return c.JSON(httpcode, map[string]interface{}{
				"message": FailedToDeleteDataActivities,
				"error":   err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  SuccessMsg,
			"message": SuccessMsg,
		})
	}
}
