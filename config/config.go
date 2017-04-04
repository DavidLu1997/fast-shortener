package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

const durationStep = time.Second

type Configuration struct {
	Server *ServerConfiguration `json:"server"`
	Links  *LinksConfiguration  `json:"links"`
	Cache  *CacheConfiguration  `json:"cache"`
}

type ServerConfiguration struct {
	BaseUrl string `json:"baseUrl"`
	Port    string `json:"port"`
}

type LinksConfiguration struct {
	MinLength   int           `json:"minLength"`
	MaxLength   int           `json:"maxLength"`
	MinDuration time.Duration `json:"minDuration"`
	MaxDuration time.Duration `json:"maxDuration"`
}

type CacheConfiguration struct {
	DefaultDuration time.Duration `json:"defaultDuration"`
	DefaultPurge    time.Duration `json:"defaultPurge"`
}

func GetConfig(filePath string) (*Configuration, error) {
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
	c.Cache.DefaultDuration *= durationStep
	c.Cache.DefaultPurge *= durationStep
}
