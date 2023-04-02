UPDATE challenge_2_be.todos
SET 
    title = ?,
    priority = ?,
    is_active = ?,
    updated_at = ?
WHERE todo_id = ?;