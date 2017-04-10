package model

import (
	"testing"
	"time"

	"github.com/davidlu1997/fast-shortener/config"
)

const configPath = "../config/config.json"

func TestPreSave(t *testing.T) {
	link := Link{
		URL:     "https://google.com",
		Key:     "a",
		Seconds: 30,
	}

	link.PreSave()

	if *link.Duration != 30*time.Second {
		t.Fatal("failed to get correct count")
	}
}

func TestIsValid(t *testing.T) {
	config, err := config.GetConfig(configPath)
	if config == nil {
		t.Fatal(err)
	}

	link1 := Link{
		URL:     "https://google.com",
		Key:     "a",
		Seconds: 300,
	}
	link1.PreSave()

	if link1.IsValid(config) {
		t.Fatal("should be invalid, key too short")
	}

	if !link1.IsValid(nil) {
		t.Fatal("should be valid with nil config")
	}

	link2 := Link{
		URL:     "https://google.com",
		Key:     "derp-herp",
		Seconds: 1,
	}
	link2.PreSave()

	if link2.IsValid(config) {
		t.Fatal("should be invalid, duration too short")
	}

	link3 := Link{
		URL:     "https://google.com",
		Key:     "derp-herp",
		Seconds: 300,
	}
	link3.PreSave()

	if !link3.IsValid(config) {
		t.Fatal("should be valid")
	}

	link4 := Link{
		URL:     "https://google.com",
		Key:     "abcdefghijklmnopqrstuvwxyz",
		Seconds: 300,
	}
	link4.PreSave()

	if link4.IsValid(config) {
		t.Fatal("should be invalid, key too long")
	}

	link5 := Link{
		URL:     "https://google.com",
		Key:     "derp-herp",
		Seconds: 3000000,
	}
	link5.PreSave()

	if link5.IsValid(config) {
		t.Fatal("should be invalid, duration too long")
	}

	link6 := Link{
		URL: "https://google.com",
		Key: "derp-herp",
	}

	if link6.IsValid(config) {
		t.Fatal("should be invalid, no duration")
	}
}
