package plugins

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const volumeIsActive = true
const volumeCheckInterval = time.Second * 120

func init() {
	List = append(List, plugin{
		Getter: func() (string, error) {
			cmd := exec.Command("pactl", "get-sink-volume", "@DEFAULT_SINK@")
			var out bytes.Buffer
			cmd.Stdout = &out
			err := cmd.Run()
			if err != nil {
				return "", err
			}

			// Example output: "Volume: front-left: 65536 / 100% / -0.00 dB, front-right: 65536 / 100% / -0.00 dB"
			output := out.String()
			lines := strings.Split(output, "\n")
			if len(lines) == 0 {
				return "", fmt.Errorf("unexpected output from pactl")
			}

			// Find the volume percentage
			volumeLine := lines[0]
			parts := strings.Fields(volumeLine)
			if len(parts) < 5 {
				return "", fmt.Errorf("unexpected format of pactl output")
			}

			// Extract the volume percentage
			volumeStr := strings.Trim(parts[4], "%")
			volume, err := strconv.Atoi(volumeStr)
			if err != nil {
				return "", err
			}

			return fmt.Sprintf(" %d%%", volume), nil
		},
		Trigger: volumeCheckInterval,
		Active:  volumeIsActive,
	})
}
