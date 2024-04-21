package config

import "github.com/go-zoox/core-utils/regexp"

// IsNotFoundErr returns true if the error is a not found error.
func IsNotFoundErr(err error) bool {
	return regexp.Match("not found", err.Error())
}
