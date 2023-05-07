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

const DefaultFileType = "YAML"

// LoadOptions is the options for Load
type LoadOptions struct {
	// FilePath is the config file path.
	FilePath string

	// options: YAML | JSON | TOML | INI | HOST, default: YAML
	Type string

	// Unique AppName for the config file, default: ""
	AppName string

	// Config name, default: config.yml, and type is YAML
	Name string
}

// Load loads the config from the given file path.
// If the file path is empty, it will load the config from the default file path.
// Default file path is
//
//	 custom file path
//		> ${PWD}/.{APP_NAME}.{NAME}.yml
//		/etc/{APP_NAME}/{NAME}.yml (user is root) | $HOME/.config/{APP_NAME}/{NAME}.yml.
//		> ${PWD}/.{APP_NAME}.yml
//		> /etc/{APP_NAME}/config.yml (user is root) | $HOME/.config/{APP_NAME}/config.yml.
//		> ${PWD}/.config.yml
func Load(config any, options ...*LoadOptions) error {
	fileType := DefaultFileType
	filepathX := fs.JoinCurrentDir(".config.yml")

	if len(options) > 0 && options[0] != nil {
		optionsX := options[0]

		if optionsX.FilePath != "" {
			filepathX = optionsX.FilePath
		} else if optionsX.AppName != "" {
			if optionsX.Name != "" {
				if fs.IsExist(fs.JoinCurrentDir(fmt.Sprintf(".%s.%s.yml", optionsX.AppName, optionsX.Name))) {
					filepathX = fs.JoinCurrentDir(fmt.Sprintf(".%s.%s.yml", optionsX.AppName, optionsX.Name))
				} else {
					filepathX = fs.JoinConfigDir(optionsX.AppName, optionsX.Name)
				}
			} else {
				// ${PWD}/.{APP_NAME}.yml > ${PWD}/.{APP_NAME}.yml > /etc/{APP_NAME}/config.yml (user is root) | $HOME/.config/{APP_NAME}/config.yml
				if fs.IsExist(fs.JoinCurrentDir(fmt.Sprintf(".%s.yml", optionsX.AppName))) {
					filepathX = fs.JoinCurrentDir(fmt.Sprintf(".%s.yml", optionsX.AppName))
				} else {
					filepathX = fs.JoinConfigDir(optionsX.AppName)
				}
			}
		}

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
