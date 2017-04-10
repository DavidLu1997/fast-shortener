package shortener

import (
	"fmt"
	"testing"

	"github.com/davidlu1997/fast-shortener/config"
	"github.com/davidlu1997/fast-shortener/model"
)

func BenchmarkCacheShortenerPut(b *testing.B) {
	config, _ := config.GetConfig(configPath)

	config.Cache.MaxSize = 2 * b.N
	shortener := InitCacheShortener(config)
	links := buildLinks(b.N)
	b.ResetTimer()

	for _, link := range links {
		shortener.Put(&link)
	}
}

func BenchmarkCacheShortenerGet(b *testing.B) {
	config, _ := config.GetConfig(configPath)

	config.Cache.MaxSize = 2 * b.N
	shortener := InitCacheShortener(config)
	links := buildLinks(b.N)

	for _, link := range links {
		shortener.Put(&link)
	}

	b.ResetTimer()

	for _, link := range links {
		shortener.Get(link.Key)
	}
}

func buildLinks(n int) []model.Link {
	links := make([]model.Link, n)
	for i := range links {
		links[i] = model.Link{
			URL:     fmt.Sprintf("http://derp.com/get-%d", i),
			Key:     fmt.Sprintf("key-%d", i),
			Seconds: 30,
		}
	}

	return links
}
