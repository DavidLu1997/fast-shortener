package model

import "github.com/davidlu1997/fast-shortener/config"
import "time"

type Link struct {
	URL       string    `json:"url"`
	Key       string    `json:"key"`
	CreatedAt time.Time `json:"-"`
	ExpiresAt time.Time `json:"-"`
}

func (l *Link) IsValid(config *config.Configuration) bool {
	if config == nil {
		return true
	}

	duration := l.ExpiresAt.Sub(l.CreatedAt)
	if duration > config.Links.MaxDuration || duration < config.Links.MinDuration {
		return false
	}

	length := len(l.Key)
	if length > config.Links.MaxLength || length < config.Links.MinLength {
		return false
	}

	return true
}
