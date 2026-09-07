//go:build wasm

package pluginwatcher

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"k8s.io/kubernetes/pkg/kubelet/util"
)

func getStat(event fsnotify.Event) (os.FileInfo, error) {
	return nil, fmt.Errorf("not supported on wasm")
}

var getSocketPath = util.NormalizePath
