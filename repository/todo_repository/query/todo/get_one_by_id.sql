SELECT
    td.todo_id "id",
    td.activity_group_id "activity_group_id",
    td.title "title",
    td.is_active "is_active",
    td.priority "priority",
    td.created_at "created_at",
    td.updated_at "updated_at",
    td.deleted_at "deleted_at"
FROM challenge_2_be.todos td
WHERE td.todo_id = ?
LIMIT 1;