package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DatabaseUser     string `split_words:"true" default:"postgres"`
	DatabasePassword string `split_words:"true" default:"postgres"`
	DatabaseAddr     string `split_words:"true" default:"localhost:5432"`
	DatabaseName     string `split_words:"true" default:"howto_dev"`
}

func (cfg *Config) Setup() error {
	err := envconfig.Process("howto", cfg)
	if err != nil {
		return err
	}
	return nil
}
