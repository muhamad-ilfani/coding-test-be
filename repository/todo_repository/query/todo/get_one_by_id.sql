SELECT
    td.todo_id "id",
    td.activity_group_id "activity_group_id",
    td.title "title",
    td.is_active "is_active",
    td.priority "priority",
    td.created_at "created_at",
    td.updated_at "updated_at"
FROM project2.todos td
WHERE td.todo_id = $1
LIMIT 1;