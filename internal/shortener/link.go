package shortener

type Link struct {
	ShortUrl  string `json:"short_url"`
	LongUrl   string `json:"long_url"`
	ExpiresAt int64  `json:"expires_at"`
}
