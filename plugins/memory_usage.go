package plugins

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const memInfoFile = "/proc/meminfo"
const memUsageIsActive = true

func init() {
	PList = append(PList, Plugin{
		Getter: func() (string, error) {
			data, err := os.ReadFile(memInfoFile)
			if err != nil {
				return "", err
			}
			lines := strings.Split(string(data), "\n")
			var totalMem, freeMem, buffers, cached int
			for _, line := range lines {
				fields := strings.Fields(line)
				if len(fields) < 2 {
					continue
				}
				value, err := strconv.Atoi(fields[1])
				if err != nil {
					continue
				}
				switch fields[0] {
				case "MemTotal:":
					totalMem = value
				case "MemFree:":
					freeMem = value
				case "Buffers:":
					buffers = value
				case "Cached:":
					cached = value
				}
			}
			usedMem := totalMem - freeMem - buffers - cached
			usedMemGB := float64(usedMem) / 1024 / 1024
			return fmt.Sprintf("▦  %.1fG", usedMemGB), nil
		},
		Span:   time.Second * 2,
		Active: memUsageIsActive,
		Order:  7,
	})
}
