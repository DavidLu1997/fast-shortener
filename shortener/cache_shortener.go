package shortener

import (
	"github.com/davidlu1997/fast-shortener/config"
	"github.com/davidlu1997/fast-shortener/model"
	cache "github.com/patrickmn/go-cache"
)

type CacheShortener struct {
	cache  *cache.Cache
	config *config.Configuration
	Shortener
}

func InitCacheShortener(config *config.Configuration) *CacheShortener {
	cache := cache.New(config.Cache.DefaultDuration, config.Cache.DefaultPurge)
	if cache == nil {
		return nil
	}

	return &CacheShortener{
		cache:  cache,
		config: config,
	}
}

func (c *CacheShortener) Get(key string) *model.Link {
	if link, found := c.cache.Get(key); found {
		return link.(*model.Link)
	}
	return nil
}

func (c *CacheShortener) Put(link *model.Link) error {
	if link == nil || !link.IsValid(c.config) {
		return nil
	}

	return c.cache.Add(link.Key, link.URL, link.Duration)
}
