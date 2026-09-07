//go:build wasm

package securitycontext

import (
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

// possibleCPUs returns the number of possible CPUs on this host.
func possibleCPUs() (cpus []int) {
	if ncpu := possibleCPUsParsed(); ncpu != nil {
		return ncpu
	}

	for i := range runtime.NumCPU() {
		cpus = append(cpus, i)
	}

	return cpus
}

// possibleCPUsParsed is parsing the amount of possible CPUs on this host from
// /sys/devices.
var possibleCPUsParsed = sync.OnceValue(func() (cpus []int) {
	data, err := os.ReadFile("/sys/devices/system/cpu/possible")
	if err != nil {
		return nil
	}

	ranges := strings.SplitSeq(strings.TrimSpace(string(data)), ",")

	for r := range ranges {
		if rStart, rEnd, ok := strings.Cut(r, "-"); !ok {
			cpu, err := strconv.Atoi(rStart)
			if err != nil {
				return nil
			}
			cpus = append(cpus, cpu)
		} else {
			var start, end int
			start, err := strconv.Atoi(rStart)
			if err != nil {
				return nil
			}
			end, err = strconv.Atoi(rEnd)
			if err != nil {
				return nil
			}
			for i := start; i <= end; i++ {
				cpus = append(cpus, i)
			}
		}
	}

	return cpus
})
