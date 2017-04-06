package shortener

import (
	"testing"
	"time"

	"github.com/davidlu1997/fast-shortener/config"
	"github.com/davidlu1997/fast-shortener/model"
)

const configPath = "../config/config.json"

func TestCacheShortenerNormal(t *testing.T) {
	config, err := config.GetConfig(configPath)
	if config == nil {
		t.Fatal(err)
	}

	cache := InitCacheShortener(config)
	if cache == nil {
		t.Fatal("failed to init cache shortener")
	}

	link1 := &model.Link{
		URL:     "https://google.com",
		Key:     "derp-herp",
		Seconds: 300,
	}

	if err := cache.Put(link1); err != nil {
		t.Fatal(err)
	}

	if l := cache.Get(link1.Key); l == nil {
		t.Fatal("should have found link")
	} else if l.URL != link1.URL || l.Key != link1.Key {
		t.Fatal("not correct link")
	}

	if l := cache.Get("abc"); l != nil {
		t.Fatal("should not have found random key")
	}
}

func TestCacheShortenerExpiration(t *testing.T) {
	config, err := config.GetConfig(configPath)
	if config == nil {
		t.Fatal(err)
	}

	config.Links.MinDuration = 1 * time.Microsecond
	config.Cache.DefaultPurge = 1 * time.Microsecond

	cache := InitCacheShortener(config)
	if cache == nil {
		t.Fatal("failed to init cache shortener")
	}

	link1 := &model.Link{
		URL: "https://google.com",
		Key: "derp-herp",
	}
	link1.Duration = new(time.Duration)
	*link1.Duration = 1 * time.Microsecond

	if err := cache.Put(link1); err != nil {
		t.Fatal(err)
	}

	time.Sleep(5 * time.Microsecond)

	if l := cache.Get(link1.Key); l != nil {
		t.Fatal("should not have found link")
	}
}

func TestCacheShortenerMaxSize(t *testing.T) {
	config, err := config.GetConfig(configPath)
	if config == nil {
		t.Fatal(err)
	}

	config.Cache.MaxSize = 1

	cache := InitCacheShortener(config)
	if cache == nil {
		t.Fatal("failed to init cache shortener")
	}

	link1 := &model.Link{
		URL:     "https://google.com",
		Key:     "derp-herp",
		Seconds: 300,
	}

	if err := cache.Put(link1); err != nil {
		t.Fatal(err)
	}

	link2 := &model.Link{
		URL:     "https://google.com",
		Key:     "ferp-gerp",
		Seconds: 300,
	}

	if err := cache.Put(link2); err == nil {
		t.Fatal("should have errored")
	}
}
