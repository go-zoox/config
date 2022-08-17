package config

import "fmt"

// Async loads the config from the given service.
func Async(config interface{}, fn func() (string, error)) error {
	raw, err := fn()
	if err != nil {
		return fmt.Errorf("load config async error: %s", err)
	}

	return Parse([]byte(raw), config)
}
