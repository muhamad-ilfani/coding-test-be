UPDATE project2.todos
SET 
    title = $2,
    priority = $3,
    is_active = $4,
    updated_at = $5
WHERE todo_id = $1
RETURNING todo_id, activity_group_id, title, is_active, priority, created_at, updated_at;