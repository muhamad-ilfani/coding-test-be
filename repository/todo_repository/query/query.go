package query

import _ "embed"

var (
	//go:embed todo/create_todo.sql
	CreateTodo string
	//go:embed todo/get_all_todos.sql
	GetAllTodos string
	//go:embed todo/get_one_by_id.sql
	GetOneTodoByID string
	//go:embed todo/update_one_by_id.sql
	UpdateOneTodoByID string
	//go:embed todo/delete_one_by_id.sql
	DeleteOneTodoByID string
)
