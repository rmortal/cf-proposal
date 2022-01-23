-- name: FindRedirectByShortUrl :one
SELECT * FROM URL
WHERE short_url = $1;

-- name: CreateUrl :one
INSERT INTO URL (long_url, short_url, expiration_dt)
VALUES($1, $2, $3) RETURNING *;