package shortener

import "time"

type Link struct {
	ID        string     `redis:"id"`
	ShortCode string     `redis:"short_code"`
	LongUrl   string     `redis:"long_url"`
	CreatedAt *time.Time `redis:"created_at"`
	ExpiresAt *time.Time `redis:"expires_at"`
}
