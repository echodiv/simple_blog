package apiserver

import "github.com/echodiv/simple_blog/internal/app/store"

// Canfig structure
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

// Create new config intance
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8000",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
