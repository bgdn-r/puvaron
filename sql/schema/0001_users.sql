-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY UNIQUE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    email VARCHAR(150) UNIQUE NOT NULL,
    phone VARCHAR(60) UNIQUE,
    first_name VARCHAR(60) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(200) NOT NULL
);

-- +goose Down
DROP TABLE users;
