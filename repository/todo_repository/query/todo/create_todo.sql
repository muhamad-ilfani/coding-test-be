INSERT INTO project2.todos(title, activity_group_id, priority, is_active, created_at, updated_at)
VALUES($1, $2, $3, $4, $5, $6)
RETURNING todo_id;