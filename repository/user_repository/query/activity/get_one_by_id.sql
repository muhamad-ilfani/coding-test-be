SELECT
    ac.activity_id "id",
    ac.title "title",
    ac.email "email",
    ac.created_at "created_at",
    ac.updated_at "updated_at"
FROM project2.activities ac
WHERE ac.activity_id = $1
LIMIT 1;