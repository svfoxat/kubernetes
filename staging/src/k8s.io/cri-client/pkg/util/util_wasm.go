//go:build js && wasm

package util

import (
	"context"
	"fmt"
	"net"
	"net/url"
)

var dialContext func(ctx context.Context, addr string) (net.Conn, error)
var listen func(addr string) (net.Listener, error)

func SetDialContext(dc func(ctx context.Context, addr string) (net.Conn, error)) {
	dialContext = dc
}

func SetListen(l func(addr string) (net.Listener, error)) {
	listen = l
}

// CreateListener creates a listener on the specified endpoint.
func CreateListener(endpoint string) (net.Listener, error) {
	protocol, addr, err := parseEndpoint(endpoint)
	if err != nil {
		return nil, err
	}
	if protocol != "vnet" && protocol != "tcp" {
		return nil, fmt.Errorf("protocol %q not supported in this build", protocol)
	}
	if listen == nil {
		return nil, fmt.Errorf("no listener installed: call util.SetListen first")
	}
	return listen(addr)
}

// GetAddressAndDialer returns the address parsed from the given endpoint and a context dialer.
func GetAddressAndDialer(endpoint string) (string, func(ctx context.Context, addr string) (net.Conn, error), error) {
	protocol, addr, err := parseEndpoint(endpoint)
	if err != nil {
		return "", nil, err
	}
	if protocol != "vnet" && protocol != "tcp" {
		return "", nil, fmt.Errorf("protocol %q not supported in this build", protocol)
	}
	return addr, dial, nil
}

func dial(ctx context.Context, addr string) (net.Conn, error) {
	if dialContext == nil {
		return nil, fmt.Errorf("no dialer installed: call util.SetDialContext first")
	}
	return dialContext(ctx, addr)
}

func parseEndpoint(endpoint string) (string, string, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return "", "", err
	}
	switch u.Scheme {
	case "vnet":
		return "vnet", u.Host, nil
	case "tcp":
		return "tcp", u.Host, nil
	case "":
		return "", "", fmt.Errorf("using %q as endpoint is deprecated, please consider using full url format", endpoint)
	default:
		return u.Scheme, "", fmt.Errorf("protocol %q not supported", u.Scheme)
	}
}
