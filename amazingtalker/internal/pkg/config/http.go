package config

import "fmt"

// HTTP declare http configuration
type HTTP struct {
	Host           string `json:"host" yaml:"host"`
	Port           int    `json:"port" yaml:"port"`
	Mode           string `json:"mode" yaml:"mode"`
	CacheExpireMin int    `json:"cache_expire_min" yaml:"cacheExpireMin"`
}

// GetAddress serve caller to get combine host and port, format is `host:port`
func (h *HTTP) GetAddress() string {
	return fmt.Sprintf("%v:%v", h.Host, h.Port)
}
