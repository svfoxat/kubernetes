//go:build js && wasm

package config

import "k8s.io/klog/v2"

func (s *sourceFile) startWatch(logger klog.Logger) {
	// fsnotify is unavailable on wasm; changes are picked up by the
	// periodic listConfig() poll in run().
	logger.V(4).Info("file watching via fsnotify unavailable on wasm; using polling")
}

func (s *sourceFile) consumeWatchEvent(logger klog.Logger, e *watchEvent) error {
	return nil // never called on wasm: startWatch emits no events
}
