package delivery

import (
	"context"

	"github.com/labstack/echo"
)

func (x *echoObject) initRoute(ctx context.Context) {
	x.Echo = echo.New()

	x.Echo.GET("/", welcome(ctx))
	x.Echo.GET("/activity-groups", GetAllActivities(ctx, x.ActivityUseCase))
	x.Echo.POST("/activity-groups", CreateActivity(ctx, x.ActivityUseCase))
	x.Echo.GET("/activity-groups/:id", GetOneByID(ctx, x.ActivityUseCase))
	x.Echo.PATCH("/activity-groups/:id", UpdateOneActivityByID(ctx, x.ActivityUseCase))
	x.Echo.DELETE("/activity-groups/:id", DeleteOneActivityByID(ctx, x.ActivityUseCase))
	x.Echo.POST("/todo-items", CreateTodo(ctx, x.TodoUseCase))
	x.Echo.GET("/todo-items", GetAllTodos(ctx, x.TodoUseCase))
	x.Echo.GET("/todo-items/:id", GetOneTodoByID(ctx, x.TodoUseCase))
	x.Echo.PATCH("/todo-items/:id", UpdateOneTodoByID(ctx, x.TodoUseCase))
	x.Echo.DELETE("/todo-items/:id", DeleteOneTodoByID(ctx, x.TodoUseCase))
}
