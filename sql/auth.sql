-- name: create-user-table
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
    username VARCHAR(255),
    password VARCHAR(255),
    oAuthCode VARCHAR(255) DEFAULT '',
    level INTEGER,
    points FLOAT,
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


-- name: set-lastChecked
UPDATE users SET lastChecked = ? AND level = ? WHERE id = ?

-- name: update-points
UPDATE user SET points = ?  WHERE id = ?