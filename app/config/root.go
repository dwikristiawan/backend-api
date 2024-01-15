package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Root struct {
	Server Server
}

func Load(filenames ...string) *Root {
	_ = godotenv.Overload(filenames...)
	r := Root{Server: Server{}}

	mustLoad("SERVER", &r.Server)
	return &r
}

func mustLoad(prefix string, spec interface{}) {
	err := envconfig.Process(prefix, spec)
	if err != nil {
		panic(err)
	}
}
