package config

import "time"

type Server struct {
	ServiceName    string        `envconfig:"SERVER_NAME" default:"awokawok"`
	Mode           string        `envconfig:"SERVER_MODE"`
	HostServer     string        `envconfig:"SERVER_HOST"`
	PortServer     string        `envconfig:"SERVER_PORT"`
	ContextTimeout time.Duration `envconfig:"SERVER_TIMEOUT" default:"2s"`
}
