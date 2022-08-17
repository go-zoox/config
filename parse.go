package config

import "github.com/go-zoox/encoding/yaml"

// Parse parses the config from the given raw config.
func Parse(data []byte, config interface{}) error {
	return yaml.Decode(data, config)
}
