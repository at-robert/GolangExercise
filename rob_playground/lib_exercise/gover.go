package gover

import (
	"runtime"
)

// GetGoVer to return the current Go Version number
func GetGoVer() string {
	return runtime.Version()
}
