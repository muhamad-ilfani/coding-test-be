INSERT INTO project2.activities(title, email, created_at, updated_at)
VALUES($1, $2, $3, $4)
RETURNING activity_id;