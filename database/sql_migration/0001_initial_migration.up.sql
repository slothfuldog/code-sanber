

-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS categories (
    id INTEGER PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(100),
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS books (
    id INTEGER PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    category_id INTEGER REFERENCES categories(id),
    description TEXT,
    image_url VARCHAR(255),
    release_year INTEGER,
    price INTEGER,
    total_page INTEGER,
    thickness VARCHAR(10),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(100),
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(100),
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by VARCHAR(100)
);

-- +migrate StatementEnd

