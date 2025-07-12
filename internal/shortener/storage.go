package shortener

import (
	"context"
	"fmt"
	"log"
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
	err := s.Cache.HGetAll(context.Background(), fmt.Sprintf("link:%s", code)).Scan(&link)

	if link.ID == "" {
		err = s.DBPool.QueryRow(context.Background(), `SELECT * FROM links WHERE short_code = $1`, code).Scan(
			&link.ID, &link.ShortCode, &link.LongUrl, &link.CreatedAt, &link.ExpiresAt)

		if err != nil {
			log.Printf("failed to query links table: %v", err)
			return nil, err
		}

		err = s.Cache.HSet(context.Background(), fmt.Sprintf("link:%s", code), link).Err()
		if err != nil {
			log.Printf("failed to set link in cache: %v", err)
			return nil, err
		}

		return &link, err
	}

	return &link, nil
}

func (s *LinkStorage) insertLink(code *string, url string, expiration *time.Time) error {
	_, err := s.DBPool.Exec(context.Background(), `INSERT INTO links (short_code, long_url, expires_at) VALUES ($1, $2, $3)`, code, url, expiration)

	return err
}
