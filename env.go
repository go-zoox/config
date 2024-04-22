package config

import (
	"github.com/go-zoox/tag"
	"github.com/go-zoox/tag/datasource"
)

// Env loads the config from the environment variables.
// The config must be a pointer to a struct.
// Example:
//
//	type Config struct {
//		Host string `env:"HOST"`
//		Port int `env:"PORT"`
//	}
//
//	var config Config
//	if err := config.Env(&config); err != nil {
//		log.Fatal(err)
//	}
func Env(config interface{}) error {
	tg := tag.New("env", datasource.NewEnvSource())

	return tg.Decode(config)
}
