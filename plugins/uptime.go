package plugins

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const uptimeIsActive = true

func init() {
	List = append(List, Plugin{
		Getter: func() (string, error) {
			data, err := os.ReadFile("/proc/uptime")
			if err != nil {
				return "", err
			}

			fields := strings.Fields(string(data))
			if len(fields) == 0 {
				return "", fmt.Errorf("unexpected content in /proc/uptime")
			}

			uptimeSeconds, err := strconv.ParseFloat(fields[0], 64)
			if err != nil {
				return "", err
			}

			uptimeDuration := time.Duration(uptimeSeconds) * time.Second

			hours := int(uptimeDuration.Hours())
			minutes := int(uptimeDuration.Minutes()) % 60
			seconds := int(uptimeDuration.Seconds()) % 60
			uptime := fmt.Sprintf("  %02d:%02d:%02d", hours, minutes, seconds)

			return uptime, nil
		},
		Span:   3 * time.Second,
		Active: uptimeIsActive,
		Order:  2,
	})
}
