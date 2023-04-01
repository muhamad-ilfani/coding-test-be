UPDATE project2.activities
SET 
    title = $2,
    updated_at = $3
WHERE activity_id = $1
RETURNING email, created_at;