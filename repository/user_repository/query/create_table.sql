CREATE TABLE IF NOT EXISTS project2.activities(
    activity_id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    created_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
    updated_at timestamp NOT NULL DEFAULT now(),
    deleted_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
    deleted_at timestamp NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS project2.todos(
    todo_id BIGSERIAL PRIMARY KEY,
    activity_group_id int8 NOT NULL,
    title VARCHAR(255) NOT NULL,
    priority VARCHAR(255) NOT NULL,
    created_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
    updated_at timestamp NOT NULL DEFAULT now(),
    deleted_by VARCHAR(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
    deleted_at timestamp NOT NULL DEFAULT now(),
    CONSTRAINT user_fk_products FOREIGN KEY (activity_group_id) REFERENCES project2.activities (activity_id)
);