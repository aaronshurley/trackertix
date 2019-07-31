package config

import (
	"errors"
	"os"
)

type Config struct {
	TrackerToken string
}

func NewConfig() (Config, error) {
	c := Config{}
	if err := c.validate(); err != nil {
		return c, err
	}

	return c, nil
}

func (c *Config) validate() error {
	c.TrackerToken = os.Getenv("PIVOTAL_TRACKER_API_TOKEN")
	if c.TrackerToken == "" {
		return errors.New("PIVOTAL_TRACKER_API_TOKEN is not set")
	}

	return nil
}
