SELECT
    td.todo_id "id"
FROM challenge_2_be.todos td
ORDER BY td.created_at DESC
LIMIT 1;