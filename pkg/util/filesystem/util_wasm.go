//go:build js && wasm

package filesystem

import (
	"fmt"
	"os"
	"path/filepath"
)

// IsUnixDomainSocket returns whether a given file is a AF_UNIX socket file
func IsUnixDomainSocket(filePath string) (bool, error) {
	return false, fmt.Errorf("IsUnixDomainSocket is not supported on WASM")
}

// Chmod is the same as os.Chmod on Unix.
func Chmod(name string, mode os.FileMode) error {
	return os.Chmod(name, mode)
}

// MkdirAll is same as os.MkdirAll on Unix.
func MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

// IsAbs is same as filepath.IsAbs on Unix.
func IsAbs(path string) bool {
	return filepath.IsAbs(path)
}
