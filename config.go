package config

import (
	"fmt"

	"github.com/go-zoox/fs"
	"github.com/go-zoox/fs/type/ini"
	"github.com/go-zoox/fs/type/json"
	"github.com/go-zoox/fs/type/toml"
	"github.com/go-zoox/fs/type/yaml"
	"github.com/go-zoox/tag"
	"github.com/go-zoox/tag/datasource"
)

// LoadOptions is the options for Load
type LoadOptions struct {
	// FilePath is the config file path.
	FilePath string

	// options: YAML | JSON | TOML | INI | HOST, default: YAML
	Type string

	// Unique ID for the config file, default: config, and type is YAML
	ID string
}

// Load loads the config from the given file path.
// If the file path is empty, it will load the config from the default file path.
// Default file path is "${PWD}/.config.yml".
func Load(config any, options ...*LoadOptions) error {
	filepathX := fs.JoinPath(fs.CurrentDir(), ".config.yml")
	fileType := "YAML"
	if len(options) > 0 {
		optionsX := options[0]
		if optionsX.FilePath != "" {
			filepathX = optionsX.FilePath
			ext := fs.ExtName(filepathX)

			switch ext {
			case "", ".yml", ".yaml":
				fileType = "YAML"
			case ".json":
				fileType = "JSON"
			case ".toml":
				fileType = "TOML"
			case ".ini":
				fileType = "INI"
			case ".host":
				fileType = "HOST"
			default:
				return fmt.Errorf("unsupported file type: %s", ext)
			}
		} else if optionsX.ID != "" {
			filepathX = fs.JoinPath(fs.CurrentDir(), "."+optionsX.ID+".yml")
		}
	}

	if !fs.IsExist(filepathX) {
		return fmt.Errorf("config path (%s) not found", filepathX)
	}

	configDataSource := make(map[string]any)
	switch fileType {
	case "YAML":
		if err := yaml.Read(filepathX, &configDataSource); err != nil {
			return err
		}
	case "JSON":
		if err := json.Read(filepathX, &configDataSource); err != nil {
			return err
		}
	case "TOML":
		if err := toml.Read(filepathX, &configDataSource); err != nil {
			return err
		}
	case "INI":
		if err := ini.Read(filepathX, &configDataSource); err != nil {
			return err
		}
	// case "HOST":
	// 	if err := hosts.Read(filepathX, &configDataSource); err != nil {
	// 		return err
	// 	}
	default:
		return fmt.Errorf("unsupported file type: %s", fileType)
	}

	tg := tag.New("config", datasource.NewMapDataSource(configDataSource))

	return tg.Decode(config)
}
