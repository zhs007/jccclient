package jccclient

import "strings"

// IsNoRetryError - is noretry
func IsNoRetryError(err error) bool {
	return strings.Index(err.Error(), "noretry:") >= 0
}
