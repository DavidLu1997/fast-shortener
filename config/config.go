package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

const durationStep = time.Second
const filePath = "config/config.json"

type Configuration struct {
	Server ServerConfiguration `json:"server"`
	Links  LinksConfiguration  `json:"links"`
}

type ServerConfiguration struct {
	BaseUrl string `json:"baseUrl"`
	Port    string `json:"port"`
}

type LinksConfiguration struct {
	MinLength   uint          `json:"minLength"`
	MaxLength   uint          `json:"maxLength"`
	MinDuration time.Duration `json:"minDuration"`
	MaxDuration time.Duration `json:"maxDuration"`
}

func GetConfig() (*Configuration, error) {
	path, _ := filepath.Abs(filePath)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	config := Configuration{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	config.processConfig()

	return &config, nil
}

func (c *Configuration) processConfig() {
	c.Links.MinDuration *= durationStep
	c.Links.MaxDuration *= durationStep
}
