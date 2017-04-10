package config

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGetConfig(t *testing.T) {
	config, err := GetConfig("config.json")

	if config == nil {
		t.Fatal("failed to get config")
	}

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetInvalidConfig(t *testing.T) {
	config, err := GetConfig("abc.json")

	if config != nil {
		t.Fatal("Should not have gotten config")
	}

	if err == nil {
		t.Fatal("Should have errored when getting invalid config")
	}

	gibberish := []byte("{not json")
	err = ioutil.WriteFile("abc.json", gibberish, 0644)
	if err != nil {
		t.Fatalf("Error while writing invalid config: %s", err)
	}

	config, err = GetConfig("abc.json")

	if config != nil {
		t.Fatal("Should not have gotten invalid config")
	}

	if err == nil {
		t.Fatal("Should have errored getting invalid config")
	}

	if err = os.Remove("abc.json"); err != nil {
		t.Fatalf("Error while cleaning up: %s", err)
	}
}
