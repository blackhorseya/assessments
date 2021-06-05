package config

import (
	"encoding/json"
	"fmt"

	"github.com/google/wire"
	"github.com/spf13/viper"
)

// Config declare configuration for application
type Config struct {
	APP  *APP  `json:"app" yaml:"app"`
	HTTP *HTTP `json:"http" yaml:"http"`
	Log  *Log  `json:"log" yaml:"log"`
}

// NewConfig serve caller to create a Config with config file path
func NewConfig(path string) (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := new(Config)
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) String() string {
	ret, _ := json.MarshalIndent(c, "", "  ")
	return string(ret)
}

// APP declare information of application
type APP struct {
	Name string `json:"name" yaml:"name"`
}

// HTTP declare http configuration
type HTTP struct {
	Host  string `json:"host" yaml:"host"`
	Port  int    `json:"port" yaml:"port"`
	Mode  string `json:"mode" yaml:"mode"`
	Limit int    `json:"limit" yaml:"limit"`
	Burst int    `json:"burst" yaml:"burst"`
}

// GetAddress serve caller to get combine host and port, format is `host:port`
func (h *HTTP) GetAddress() string {
	return fmt.Sprintf("%v:%v", h.Host, h.Port)
}

// Log declare log configuration
type Log struct {
	Format string `json:"format" yaml:"format"`
	Level  string `json:"level" yaml:"level"`
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewConfig)
