package model

import (
	"testing"
	"time"

	"github.com/davidlu1997/fast-shortener/config"
)

const configPath = "../config/config.json"

func TestIsValid(t *testing.T) {
	config, err := config.GetConfig(configPath)
	if config == nil {
		t.Fatal(err)
	}

	link1 := Link{
		URL:      "https://google.com",
		Key:      "a",
		Duration: 5 * time.Minute,
	}

	if link1.IsValid(config) {
		t.Fatal("should be invalid, key too short")
	}

	link2 := Link{
		URL:      "https://google.com",
		Key:      "derp-herp",
		Duration: 1 * time.Second,
	}

	if link2.IsValid(config) {
		t.Fatal("should be invalid, duration too short")
	}

	link3 := Link{
		URL:      "https://google.com",
		Key:      "derp-herp",
		Duration: 5 * time.Minute,
	}

	if !link3.IsValid(config) {
		t.Fatal("should be valid")
	}

	link4 := Link{
		URL:      "https://google.com",
		Key:      "abcdefghijklmnopqrstuvwxyz",
		Duration: 5 * time.Minute,
	}

	if link4.IsValid(config) {
		t.Fatal("should be invalid, key too long")
	}

	link5 := Link{
		URL:      "https://google.com",
		Key:      "derp-herp",
		Duration: 100 * time.Hour,
	}

	if link5.IsValid(config) {
		t.Fatal("should be invalid, duration too long")
	}
}
