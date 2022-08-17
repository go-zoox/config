package config

import (
	"github.com/go-zoox/encoding/yaml"
	"github.com/go-zoox/tag"
	"github.com/go-zoox/tag/datasource"
)

// Parse parses the config from the given raw config.
func Parse(data []byte, config interface{}) error {
	configYaml := make(map[string]any)
	if err := yaml.Decode(data, &configYaml); err != nil {
		return err
	}

	tg := tag.New("config", datasource.NewMapDataSource(configYaml))

	return tg.Decode(config)
}
