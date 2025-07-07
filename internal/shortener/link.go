package shortener

import "time"

type Link struct {
	ID        string     `json:"id"`
	ShortCode string     `json:"short_code"`
	LongUrl   string     `json:"long_url"`
	CreatedAt *time.Time `json:"created_at"`
	ExpiresAt *time.Time `json:"expires_at"`
}
