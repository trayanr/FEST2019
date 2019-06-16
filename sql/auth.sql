-- name: create-user-table
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
    username VARCHAR(255),
    email VARCHAR(255), 
    password VARCHAR(255),
    oAuthCode VARCHAR(255)
);

-- name: create-user
INSERT INTO users (username, email, password) VALUES (?, ?, ?);

-- name: get-user-by-username-passwordhash
SELECT * FROM users
WHERE username = ? AND password = ? LIMIT 1

-- name: set-user-auth
UPDATE users SET oAuthCode = ? WHERE id = ?

-- name: create-session-table