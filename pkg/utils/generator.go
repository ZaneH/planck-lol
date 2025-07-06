package pkg

import (
	"math/rand/v2"
	"strings"
)

const CODE_LENGTH = 8

func GenerateShortUrl() string {
	alphabet := "abcdefghijklmnopqrstuvwxyz0123456789"
	sb := strings.Builder{}

	for range CODE_LENGTH {
		idx := rand.IntN(len(alphabet))
		sb.WriteByte(alphabet[idx])
	}

	return sb.String()
}
