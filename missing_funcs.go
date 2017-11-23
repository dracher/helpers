package helpers

import (
	"os"
)

// FileExists check a file exists or not
func FileExists(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}
