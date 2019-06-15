-- name: create-user-table
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
    name VARCHAR(255),
    email VARCHAR(255), 
    password VARCHAR(255)
);

-- name: create-user
INSERT INTO users (name, email, password) VALUES (?, ?, ?);

-- name: get-user-by-username-passwordhash
SELECT * FROM users
WHERE name = ? AND password = ? LIMIT 1