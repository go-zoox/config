package config

import "github.com/go-zoox/core-utils/regexp"

func IsNotFoundErr(err error) bool {
	return regexp.Match("not found", err.Error())
}
