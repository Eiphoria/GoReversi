CREATE TABLE IF NOT EXISTS users (
    username VARCHAR(30) PRIMARY KEY,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);