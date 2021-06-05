package config

// Log declare log configuration
type Log struct {
	Format string `json:"format" yaml:"format"`
	Level  string `json:"level" yaml:"level"`
}
