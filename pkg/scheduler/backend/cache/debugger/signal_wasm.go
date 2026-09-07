//go:build js

package debugger

import "os"

// compareSignal is a no-op on WASM (no real signals).
var compareSignal os.Signal = os.Kill
