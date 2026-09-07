//go:build js && wasm

package cadvisor

import (
	"context"
	"os"
	"strconv"
	"strings"

	cadvisorapi "github.com/google/cadvisor/info/v1"
	cadvisorapiv2 "github.com/google/cadvisor/info/v2"
	"k8s.io/klog/v2"
)

type cadvisorWasm struct{}

var _ Interface = (*cadvisorWasm)(nil)

func New(_ klog.Logger, _ ImageFsInfoProvider, _ string, _ []string, _, _ bool) (Interface, error) {
	return &cadvisorWasm{}, nil
}

func IsPsiEnabled(_ klog.Logger) bool { return false }

func (c *cadvisorWasm) Start() error { return nil }

func (c *cadvisorWasm) MachineInfo() (*cadvisorapi.MachineInfo, error) {
	numCores := readCPUCount()
	return &cadvisorapi.MachineInfo{
		NumCores:         numCores,
		NumPhysicalCores: numCores,
		NumSockets:       1,
		MemoryCapacity:   readMemTotal(),
		MachineID:        "microkube",
		SystemUUID:       "microkube",
		BootID:           "microkube",
	}, nil
}

func (c *cadvisorWasm) VersionInfo() (*cadvisorapi.VersionInfo, error) {
	return &cadvisorapi.VersionInfo{
		KernelVersion:      "0.0.0-wasm",
		ContainerOsVersion: "WASM/JS Runtime",
	}, nil
}

func (c *cadvisorWasm) ContainerInfoV2(string, cadvisorapiv2.RequestOptions) (map[string]cadvisorapiv2.ContainerInfo, error) {
	return map[string]cadvisorapiv2.ContainerInfo{}, nil
}
func (c *cadvisorWasm) GetRequestedContainersInfo(string, cadvisorapiv2.RequestOptions) (map[string]*cadvisorapi.ContainerInfo, error) {
	return map[string]*cadvisorapi.ContainerInfo{}, nil
}
func (c *cadvisorWasm) ImagesFsInfo(context.Context) (cadvisorapiv2.FsInfo, error) {
	return cadvisorapiv2.FsInfo{}, nil
}
func (c *cadvisorWasm) RootFsInfo() (cadvisorapiv2.FsInfo, error) {
	return cadvisorapiv2.FsInfo{}, nil
}
func (c *cadvisorWasm) ContainerFsInfo(context.Context) (cadvisorapiv2.FsInfo, error) {
	return cadvisorapiv2.FsInfo{}, nil
}
func (c *cadvisorWasm) GetDirFsInfo(string) (cadvisorapiv2.FsInfo, error) {
	return cadvisorapiv2.FsInfo{}, nil
}

func readCPUCount() int {
	data, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		return 1
	}
	n := 0
	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "processor") {
			n++
		}
	}
	if n == 0 {
		return 1
	}
	return n
}

func readMemTotal() uint64 {
	data, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return 8 << 30 // 8 GiB fallback
	}
	for _, line := range strings.Split(string(data), "\n") {
		f := strings.Fields(line)
		if len(f) >= 2 && f[0] == "MemTotal:" {
			if kb, err := strconv.ParseUint(f[1], 10, 64); err == nil {
				return kb * 1024
			}
		}
	}
	return 8 << 30
}
