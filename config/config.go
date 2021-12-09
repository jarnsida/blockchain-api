package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

// Config is a config :).
type Config struct {
	HTTPAddr        string `envconfig:"HTTP_ADDR"`
	InfuraEndpoint  string `envconfig:"INFURA_ENDPOINT"`
	ContractAddress string `envconfig:"CONTRACT_ADDRESS"`
}

// Get reads config from environment. Once.
func Get() (*Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, fmt.Errorf("envconfig.Process failed: %w", err)
	}

	return &config, nil
}
