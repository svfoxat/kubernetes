//go:build js && wasm

package options

import (
	"syscall"
)

func permitPortReuse(network, addr string, conn syscall.RawConn) error {
	return nil
}

func permitAddressReuse(network, addr string, conn syscall.RawConn) error {
	return nil
}
