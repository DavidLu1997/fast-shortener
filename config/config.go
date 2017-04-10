package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

const durationStep = time.Second

// Configuration is all configuration used by fast-shortener
type Configuration struct {
	Server *ServerConfiguration `json:"server"`
	Links  *LinksConfiguration  `json:"links"`
	Cache  *CacheConfiguration  `json:"cache"`
}

// ServerConfiguration specifies the baseURL and the port
type ServerConfiguration struct {
	BaseURL string `json:"baseURL"`
	Port    string `json:"port"`
}

// LinksConfiguration specifies the min/max key lengths and duration
type LinksConfiguration struct {
	MinLength   int           `json:"minLength"`
	MaxLength   int           `json:"maxLength"`
	MinDuration time.Duration `json:"minDuration"`
	MaxDuration time.Duration `json:"maxDuration"`
}

// CacheConfiguration specifies the cache configuration
type CacheConfiguration struct {
	// DefaultDuration is unused
	DefaultDuration time.Duration `json:"defaultDuration"`

	// DefaultPurge is the maximum duration a link will be live after expiring
	DefaultPurge time.Duration `json:"defaultPurge"`

	// MaxSize is the maximum number of links in the cache
	MaxSize int `json:"maxSize"`
}

// GetConfig builds a configuration given a path to a config file
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
