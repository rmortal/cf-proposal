// Code generated by sqlc. DO NOT EDIT.
// source: URL_QUERIES.sql

package datastore

import (
	"context"
)

const createUrl = `-- name: CreateUrl :one
INSERT INTO URL (long_url, short_url, expiration_dt)
VALUES($1, $2, $3) RETURNING url_id, long_url, short_url, expiration_dt
`

type CreateUrlParams struct {
	LongUrl      string `json:"long_url"`
	ShortUrl     string `json:"short_url"`
	ExpirationDt string `json:"expiration_dt"`
}

func (q *Queries) CreateUrl(ctx context.Context, arg CreateUrlParams) (Url, error) {
	row := q.db.QueryRowContext(ctx, createUrl, arg.LongUrl, arg.ShortUrl, arg.ExpirationDt)
	var i Url
	err := row.Scan(
		&i.UrlID,
		&i.LongUrl,
		&i.ShortUrl,
		&i.ExpirationDt,
	)
	return i, err
}

const deleteUrl = `-- name: DeleteUrl :exec
DELETE FROM URL 
WHERE url_id = $1
`

func (q *Queries) DeleteUrl(ctx context.Context, urlID int32) error {
	_, err := q.db.ExecContext(ctx, deleteUrl, urlID)
	return err
}

const findRedirectByShortUrl = `-- name: FindRedirectByShortUrl :one
SELECT url_id, long_url FROM URL
WHERE short_url = $1
`

type FindRedirectByShortUrlRow struct {
	UrlID   int32  `json:"url_id"`
	LongUrl string `json:"long_url"`
}

func (q *Queries) FindRedirectByShortUrl(ctx context.Context, shortUrl string) (FindRedirectByShortUrlRow, error) {
	row := q.db.QueryRowContext(ctx, findRedirectByShortUrl, shortUrl)
	var i FindRedirectByShortUrlRow
	err := row.Scan(&i.UrlID, &i.LongUrl)
	return i, err
}

const findShortUrlByLongUrl = `-- name: FindShortUrlByLongUrl :one
SELECT url_id, long_url, short_url, expiration_dt FROM URL
WHERE long_url = $1
`

func (q *Queries) FindShortUrlByLongUrl(ctx context.Context, longUrl string) (Url, error) {
	row := q.db.QueryRowContext(ctx, findShortUrlByLongUrl, longUrl)
	var i Url
	err := row.Scan(
		&i.UrlID,
		&i.LongUrl,
		&i.ShortUrl,
		&i.ExpirationDt,
	)
	return i, err
}

const getUrl = `-- name: GetUrl :one
SELECT url_id, long_url, short_url, expiration_dt FROM URL
WHERE url_id = $1
`

func (q *Queries) GetUrl(ctx context.Context, urlID int32) (Url, error) {
	row := q.db.QueryRowContext(ctx, getUrl, urlID)
	var i Url
	err := row.Scan(
		&i.UrlID,
		&i.LongUrl,
		&i.ShortUrl,
		&i.ExpirationDt,
	)
	return i, err
}
