package todo_repo

import (
	"coding-test-be/repository"
	"coding-test-be/repository/todo_repository/query"
	"context"
	"log"
	"net/http"
	"time"
)

func (x *PostgreSQLConn) CreateTodo(
	ctx context.Context, req repository.CreateTodoRequest) (
	res repository.CreateTodoResponse, httpcode int, err error,
) {
	var id int64
	createdTime := time.Now()

	query := query.CreateTodo
	args := List{
		req.Title,
		req.ActivityGroupID,
		req.Priority,
		req.IsActive,
		createdTime,
		createdTime,
	}

	rows, err := x.tc.Query(query, args...)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Println(err)
		}
	}

	res = repository.CreateTodoResponse{
		ID:              id,
		Title:           req.Title,
		ActivityGroupID: req.ActivityGroupID,
		IsActive:        req.IsActive,
		Priority:        req.Priority,
		CreatedAt:       createdTime,
		UpdatedAt:       createdTime,
	}

	return res, httpcode, err
}

func (x *PostgreSQLConn) GetAllTodos(
	ctx context.Context, req repository.GetAllTodosRequest) (
	res repository.GetAllTodosResponse, httpcode int, err error,
) {
	query := query.GetAllTodos
	args := List{}

	rows, err := x.tc.Query(query, args...)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	defer rows.Close()

	for rows.Next() {
		data := repository.TodoList{}
		err := rows.Scan(&data.ID, &data.ActivityGroupID, &data.Title, &data.IsActive, &data.Priority, &data.CreatedAt, &data.UpdatedAt)
		if err != nil {
			log.Println(err)
		}
		res = append(res, data)
	}

	return res, httpcode, err
}

func (x *PostgreSQLConn) GetOneTodoByID(
	ctx context.Context, req repository.GetOneTodoByIDRequest) (
	res repository.GetOneTodoByIDResponse, httpcode int, err error,
) {
	query := query.GetOneTodoByID
	args := List{
		req.ID,
	}

	rows, err := x.tc.Query(query, args...)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	defer rows.Close()

	for rows.Next() {
		data := repository.GetOneTodoByIDResponse{}
		err := rows.Scan(&data.ID, &data.ActivityGroupID, &data.Title, &data.IsActive, &data.Priority, &data.CreatedAt, &data.UpdatedAt)
		if err != nil {
			log.Println(err)
		}
		res = data
	}

	return res, httpcode, err
}

func (x *PostgreSQLConn) UpdateOneTodoByID(
	ctx context.Context, req repository.UpdateOneTodoByIDRequest) (
	res repository.UpdateOneTodoByIDResponse, httpcode int, err error,
) {
	updatedTime := time.Now()

	query := query.UpdateOneTodoByID
	args := List{
		req.ID,
		req.Title,
		req.Priority,
		req.IsActive,
		updatedTime,
	}

	rows, err := x.tc.Query(query, args...)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	defer rows.Close()

	for rows.Next() {
		data := repository.UpdateOneTodoByIDResponse{}
		err := rows.Scan(&data.ID, &data.ActivityGroupID, &data.Title, &data.IsActive, &data.Priority, &data.CreatedAt, &data.UpdatedAt)
		if err != nil {
			log.Println(err)
		}
		res = data
	}

	return res, httpcode, err
}

func (x *PostgreSQLConn) DeleteOneTodoByID(
	ctx context.Context, req repository.DeleteOneTodoByIDRequest) (
	res repository.DeleteOneTodoByIDResponse, httpcode int, err error,
) {
	query := query.DeleteOneTodoByID
	args := List{
		req.ID,
	}

	_, err = x.tc.Exec(query, args...)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	res = repository.DeleteOneTodoByIDResponse{ID: req.ID}

	return res, httpcode, err
}
