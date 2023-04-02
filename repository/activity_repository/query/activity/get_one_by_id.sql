SELECT
    ac.activity_id "id",
    ac.title "title",
    ac.email "email",
    ac.created_at "created_at",
    ac.updated_at "updated_at",
    ac.deleted_at "deleted_at"
FROM challenge_2_be.activities ac
WHERE ac.activity_id = ?
LIMIT 1;