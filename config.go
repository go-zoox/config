package config

import (
	"fmt"

	"github.com/go-zoox/fs"
	"github.com/go-zoox/fs/type/yaml"
	"github.com/go-zoox/tag"
	"github.com/go-zoox/tag/datasource"
)

// LoadOptions is the options for Load
type LoadOptions struct {
	ID       string
	FilePath string
}

// Load loads the config from the given file path.
// If the file path is empty, it will load the config from the default file path.
// Default file path is "${PWD}/.config.yml".
func Load(config any, options ...*LoadOptions) error {
	filepathX := fs.JoinPath(fs.CurrentDir(), ".config.yml")
	if len(options) > 0 {
		optionsX := options[0]
		if optionsX.FilePath != "" {
			filepathX = optionsX.FilePath
		} else if optionsX.ID != "" {
			filepathX = fs.JoinPath(fs.CurrentDir(), "."+optionsX.ID+".yml")
		}
	}

	if !fs.IsExist(filepathX) {
		return fmt.Errorf("config path (%s) not found", filepathX)
	}

	configYaml := make(map[string]any)
	if err := yaml.Read(filepathX, &configYaml); err != nil {
		return err
	}

	// j, _ := json.MarshalIndent(configYaml, "", "  ")
	// fmt.Println(string(j))

	tg := tag.New("config", datasource.NewMapDataSource(configYaml))

	return tg.Decode(config)
}
