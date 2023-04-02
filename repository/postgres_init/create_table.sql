CREATE TABLE IF NOT EXISTS challenge_2_be.activities(
    activity_id bigint AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    created_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM',
    created_at timestamp NOT NULL DEFAULT now(),
    updated_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM',
    updated_at timestamp NOT NULL DEFAULT now(),
    deleted_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM',
    deleted_at timestamp NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS challenge_2_be.todos(
    todo_id bigint AUTO_INCREMENT PRIMARY KEY,
    activity_group_id int8 NOT NULL,
    title VARCHAR(255) NOT NULL,
    priority VARCHAR(255) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM',
    created_at timestamp NOT NULL DEFAULT now(),
    updated_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM',
    updated_at timestamp NOT NULL DEFAULT now(),
    deleted_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM',
    deleted_at timestamp NOT NULL DEFAULT now(),
    CONSTRAINT user_fk_products FOREIGN KEY (activity_group_id) REFERENCES challenge_2_be.activities (activity_id)
);
