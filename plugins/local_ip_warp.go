package plugins

import (
	"os/exec"
	"time"
)

const localIPWARPIsActive = true

func init() {
	List = append(List, Plugin{
		Getter: func() (string, error) {
			ip, err := getLocalIP()
			if err != nil {
				return "", err
			}
			wa, err := isServiceActive("warp-svc.service")
			if err != nil {
				return "", err
			}

			if !wa {
				return ip, nil
			}
			return ip + " +WARP", nil
		},
		Span:      time.Second * 120,
		ErrorSpan: 10 * time.Second,
		Active:    localIPWARPIsActive,
		Order:     9,
	})
}

func isServiceActive(serviceName string) (bool, error) {
	// Execute the systemctl command to check the service status
	cmd := exec.Command("systemctl", "is-active", serviceName)
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}

	// Convert the output to a string and trim any whitespace
	status := string(output)
	status = status[:len(status)-1] // Remove the trailing newline

	return status == "active", nil
}
