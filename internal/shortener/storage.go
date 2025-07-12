package shortener

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type LinkStorage struct {
	DBPool *pgxpool.Pool
	Cache  *redis.Client
}

func (s *LinkStorage) findByShortCode(code string) (*Link, error) {
	var link Link
	err := s.DBPool.QueryRow(context.Background(), `SELECT * FROM links WHERE short_code = $1`, code).Scan(
		&link.ID, &link.ShortCode, &link.LongUrl, &link.CreatedAt, &link.ExpiresAt)
	return &link, err
}

func (s *LinkStorage) insertLink(code *string, url string, expiration *time.Time) error {
	_, err := s.DBPool.Exec(context.Background(), `INSERT INTO links (short_code, long_url, expires_at) VALUES ($1, $2, $3)`, code, url, expiration)

	return err
}
