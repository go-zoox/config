package config

import (
	"github.com/go-zoox/fs"
	"github.com/go-zoox/fs/type/yaml"
	"github.com/go-zoox/tag"
	"github.com/go-zoox/tag/datasource"
)

type LoadOptions struct {
	Id       string
	FilePath string
}

func Load(config any, options ...LoadOptions) error {
	filepathX := fs.JoinPath(fs.CurrentDir(), ".config.yml")
	if len(options) > 0 {
		optionsX := options[0]
		if optionsX.FilePath != "" {
			filepathX = optionsX.FilePath
		} else if optionsX.Id != "" {
			filepathX = fs.JoinPath(fs.CurrentDir(), "."+optionsX.Id+".yml")
		}
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
