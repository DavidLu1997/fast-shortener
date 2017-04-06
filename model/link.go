package model

import "github.com/davidlu1997/fast-shortener/config"
import "time"

type Link struct {
	URL      string         `json:"url"`
	Key      string         `json:"key"`
	Duration *time.Duration `json:"-"`
	Seconds  int64          `json:"seconds"`
}

func (l *Link) PreSave() {
	if l.Duration == nil {
		l.Duration = new(time.Duration)
		*l.Duration = time.Duration(l.Seconds) * time.Second
	}
}

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
