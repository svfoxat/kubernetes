//go:build wasm

package util

import (
	"net/url"
	"path/filepath"
)

const (
	// unixProtocol is the network protocol of unix socket.
	vnetProcotol = "vnet"
)

// LocalEndpoint returns the full path to a unix socket at the given endpoint
func LocalEndpoint(path, file string) (string, error) {
	u := url.URL{
		Scheme: vnetProcotol,
		Path:   path,
	}
	return filepath.Join(u.String(), file+".sock"), nil
}

// NormalizePath is a no-op for Linux for now
func NormalizePath(path string) string {
	return path
}
