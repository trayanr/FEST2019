-- name: create-user-table
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
    username VARCHAR(255),
    password VARCHAR(255),
    oAuthCode VARCHAR(255),
    level INTEGER,
    points INTEGER,
    lastChecked INTEGER
);

-- name: create-user
INSERT INTO users (username, password) VALUES (?, ?);

-- name: get-user-by-username
SELECT * FROM users
WHERE username = ? LIMIT 1

-- name: set-user-auth
UPDATE users SET oAuthCode = ? WHERE id = ?

-- name: get-user-by-id
SELECT * FROM users
WHERE id = ?

-- name: create-session-table