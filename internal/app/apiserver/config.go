package apiserver

import "{GIT_PATH}/internal/app/helpers"

// ServerConfig - конфигурация сервера
type ServerConfig struct {
	BindAddr string
	LogLevel string
}

// NewServerConfig ...
func NewServerConfig() *ServerConfig {
	s := &ServerConfig{}
	s.BindAddr = helpers.GetEnv("BIND_ADDR", ":4000")
	s.LogLevel = helpers.GetEnv("LOG_LEVEL", "debug")

	return s
}

