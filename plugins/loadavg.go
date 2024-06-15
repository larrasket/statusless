package plugins

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const loadIsActive = true

func init() {
	List = append(List, Plugin{
		Getter: func() (string, error) {

			const loadAvgFile = "/proc/loadavg"

			data, err := os.ReadFile(loadAvgFile)
			if err != nil {
				return "", err
			}
			fields := strings.Fields(string(data))
			if len(fields) == 0 {
				return "", fmt.Errorf("unexpected content in /proc/loadavg")
			}
			loadAvgLastMinute := fields[0]
			return "  " + loadAvgLastMinute, nil
		},
		Span:   time.Second * 15,
		Active: loadIsActive,
		Order:  5,
	})
}
