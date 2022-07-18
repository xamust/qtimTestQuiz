package server

import "github.com/xamust/qtimTestQuiz/internal/app/counter"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Counter  *counter.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":9090", //default param (watch in config.toml)...
		LogLevel: "info",  //default param (watch in config.toml)...
	}
}
