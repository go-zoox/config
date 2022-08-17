package config

// LoadFromService loads the config from the given service.
func LoadFromService(config interface{}, fn func() (string, error)) error {
	return Async(config, fn)
}
