package config

import "testing"

func TestGetConfig(t *testing.T) {
	config, err := GetConfig("config.json")

	if config == nil {
		t.Fatal("failed to get config")
	}

	if err != nil {
		t.Fatal(err)
	}
}
