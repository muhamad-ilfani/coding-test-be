package repository

import (
	"context"
	"time"
)

type TodoRepo interface {
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
	GetLatestIDTodo(
		ctx context.Context, req GetLatestIDTodoRequest) (
		res GetLatestIDTodoResponse, httpcode int, err error,
	)
}

type TodoList struct {
	ID              int64
	ActivityGroupID int64
	Title           string
	IsActive        bool
	Priority        string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}

type CreateTodoRequest struct {
	Title           string
	ActivityGroupID int64
	IsActive        bool
	Priority        string
}

type CreateTodoResponse TodoList

type GetAllTodosRequest struct{}
type GetAllTodosResponse []TodoList

type GetOneTodoByIDRequest struct{ ID int64 }
type GetOneTodoByIDResponse TodoList

type UpdateOneTodoByIDRequest struct {
	ID       int64
	Title    string
	Status   string
	IsActive bool
	Priority string
}

type UpdateOneTodoByIDResponse TodoList

type DeleteOneTodoByIDRequest struct {
	ID int64
}
type DeleteOneTodoByIDResponse DeleteOneTodoByIDRequest

type GetLatestIDTodoRequest struct{}
type GetLatestIDTodoResponse struct{ ID int64 }
