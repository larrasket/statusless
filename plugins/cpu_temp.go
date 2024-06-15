package plugins

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const tempFile = "/sys/class/thermal/thermal_zone0/temp"
const cpuTempIsActive = true

func init() {
	List = append(List, Plugin{
		Getter: func() (string, error) {

			data, err := os.ReadFile(tempFile)
			if err != nil {
				return "", err
			}
			tempStr := strings.TrimSpace(string(data))

			tempMilli, err := strconv.Atoi(tempStr)
			if err != nil {
				return "", err
			}
			tempCelsius := float64(tempMilli) / 1000.0

			return fmt.Sprintf("  %.1f°C", tempCelsius), nil
		},
		Span:   time.Second * 120,
		Active: cpuTempIsActive,
		Order:  6,
	})
}
