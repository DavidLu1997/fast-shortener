package model

import "github.com/davidlu1997/fast-shortener/config"
import "time"

// Link is a link
type Link struct {
	// URL is the original unshortened URL
	URL string `json:"url"`

	// Key is the shortened key
	Key string `json:"key"`

	// Duration is the persistence time
	Duration *time.Duration `json:"-"`

	// Seconds is the persistence time in seconds
	Seconds int64 `json:"seconds"`
}

// PreSave calculates the duration given seconds
// called before the link is saved
func (l *Link) PreSave() {
	if l.Duration == nil {
		l.Duration = new(time.Duration)
		*l.Duration = time.Duration(l.Seconds) * time.Second
	}
}

// IsValid checks if a link is valid
func (l *Link) IsValid(config *config.Configuration) bool {
	if config == nil {
		return true
	}

	if l.Duration == nil {
		return false
	}

	if *l.Duration > config.Links.MaxDuration || *l.Duration < config.Links.MinDuration {
		return false
	}

	length := len(l.Key)
	if length > config.Links.MaxLength || length < config.Links.MinLength {
		return false
	}

	return true
}
