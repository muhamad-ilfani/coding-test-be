package usecases

import "context"

type TodoUseCase interface {
	CreateTodo(
		ctx context.Context, req CreateTodoRequest) (
		res CreateTodoResponse, httpcode int, err error,
	)
	GetAllTodos(
		ctx context.Context, req GetAllTodosRequest) (
		res GetAllTodosResponse, httpcode int, err error,
	)
	GetOneTodoByID(
		ctx context.Context, req GetOneTodoByIDRequest) (
		res GetOneTodoByIDResponse, httpcode int, err error,
	)
	UpdateOneTodoByID(
		ctx context.Context, req UpdateOneTodoByIDRequest) (
		res UpdateOneTodoByIDResponse, httpcode int, err error,
	)
	DeleteOneTodoByID(
		ctx context.Context, req DeleteOneTodoByIDRequest) (
		res DeleteOneTodoByIDResponse, httpcode int, err error,
	)
}

type TodoList struct {
	ID              int64  `json:"id"`
	ActivityGroupID int64  `json:"activity_group_id"`
	Title           string `json:"title"`
	IsActive        bool   `json:"is_active"`
	Priority        string `json:"priority"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
}

type CreateTodoRequest struct {
	Title           string `json:"title"`
	ActivityGroupID int64  `json:"activity_group_id"`
	IsActive        bool   `json:"is_active"`
	Priority        string `json:"priority,omitempty"`
}

type CreateTodoResponse TodoList

type GetAllTodosRequest struct{}
type GetAllTodosResponse []TodoList

type GetOneTodoByIDRequest struct{ ID int64 }
type GetOneTodoByIDResponse TodoList

type UpdateOneTodoByIDRequest struct {
	ID       int64  `json:"-"`
	Title    string `json:"title,omitempty"`
	Status   string `json:"status,omitempty"`
	IsActive bool   `json:"is_active,omitempty"`
	Priority string `json:"priority,omitempty"`
}

type UpdateOneTodoByIDResponse TodoList

type DeleteOneTodoByIDRequest struct {
	ID int64
}
type DeleteOneTodoByIDResponse DeleteOneTodoByIDRequest
