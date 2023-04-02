SELECT
    ac.activity_id "id"
FROM challenge_2_be.activities ac
ORDER BY ac.created_at DESC
LIMIT 1;