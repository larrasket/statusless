package plugins

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const batteryIsActive = true
const batteryPath = "/sys/class/power_supply/BAT0/capacity"

func init() {
	List = append(List, plugin{
		Getter: func() (string, error) {
			data, err := os.ReadFile(batteryPath)
			batteryPercentageStr := strings.TrimSpace(string(data))
			batteryPercentage, err := strconv.Atoi(batteryPercentageStr)
			if err != nil {
				return "", err
			}

			return fmt.Sprintf(" %d%%", batteryPercentage), nil
		},
		Trigger: time.Second,
		Active:  batteryIsActive,
		Name:    "battery",
	})
}
