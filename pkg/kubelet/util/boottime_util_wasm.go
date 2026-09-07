//go:build wasm

package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetBootTime() (time.Time, error) {
	bootTime, err := getBootTimeWithProcStat()
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to get boot time from /proc/stat: %w", err)
	}
	return bootTime, nil
}

func getBootTimeWithProcStat() (time.Time, error) {
	raw, err := os.ReadFile("/proc/stat")
	if err != nil {
		return time.Time{}, fmt.Errorf("error getting boot time: %w", err)
	}
	rawFields := strings.Fields(string(raw))
	for i, v := range rawFields {
		if v == "btime" {
			if len(rawFields) > i+1 {
				sec, err := strconv.ParseInt(rawFields[i+1], 10, 64)
				if err != nil {
					return time.Time{}, fmt.Errorf("error parsing boot time %s: %w", rawFields[i+1], err)
				}
				return time.Unix(sec, 0), nil
			}
			break
		}
	}

	return time.Time{}, fmt.Errorf("can not find btime from /proc/stat: %s", raw)
}
