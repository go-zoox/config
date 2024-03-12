package config

import (
	"fmt"

	"github.com/go-zoox/encoding/ini"
	"github.com/go-zoox/encoding/json"
	"github.com/go-zoox/encoding/toml"
	"github.com/go-zoox/encoding/yaml"
	"github.com/go-zoox/tag"
	"github.com/go-zoox/tag/datasource"
)

type ParseOptions struct {
	// The type of the config file, default is "YAML".
	// options: YAML | JSON | TOML | INI | HOST
	Type string
}

// Parse parses the config from the given raw config.
func Parse(data []byte, config interface{}, opts ...*ParseOptions) error {
	opt := &ParseOptions{
		Type: "YAML",
	}
	for _, o := range opts {
		if o == nil {
			continue
		}

		opt.Type = o.Type
	}

	configYaml := make(map[string]any)

	switch opt.Type {
	case "YAML":
		if err := yaml.Decode(data, &configYaml); err != nil {
			return err
		}
	case "JSON":
		if err := json.Decode(data, &configYaml); err != nil {
			return err
		}
	case "TOML":
		if err := toml.Decode(data, &configYaml); err != nil {
			return err
		}
	case "INI":
		if err := ini.Decode(data, &configYaml); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported config type: %s", opt.Type)
	}

	tg := tag.New("config", datasource.NewMapDataSource(configYaml))

	return tg.Decode(config)
}
