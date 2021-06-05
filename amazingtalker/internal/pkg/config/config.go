package config

import (
	"encoding/json"

	"github.com/google/wire"
	"github.com/spf13/viper"
)

// Config declare configuration for application
type Config struct {
	HTTP *HTTP `json:"http" yaml:"http"`
	DB   *DB   `json:"db" yaml:"db"`
	Log  *Log  `json:"log" yaml:"log"`
}

// NewConfig serve caller to create a Config with config file path
func NewConfig(path string) (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{
		HTTP: &HTTP{
			Host:           "",
			Port:           8080,
			Mode:           "debug",
			CacheExpireMin: 10,
		},
		DB: &DB{
			URL:   "",
			Debug: false,
		},
		Log: &Log{
			Format: "text",
			Level:  "debug",
		},
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) String() string {
	ret, _ := json.MarshalIndent(c, "", "  ")
	return string(ret)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewConfig)
