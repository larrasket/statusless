package plugins

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

const emacsClockIsActive = true
const emacsClockMaxLen = 28

func init() {
	PList = append(PList, Plugin{
		Getter: func() (string, error) {
			cmd := exec.Command("emacsclient", "-e", "(when (org-clocking-p) (substring-no-properties (org-clock-get-clock-string)))")

			output, err := cmd.Output()
			if err != nil {
				return "", err
			}

			result := strings.TrimSpace(string(output))
			result = strings.Trim(result, "\"")
			result = strings.NewReplacer("(", "", ")", "").Replace(result)

			if result == "nil" {
				return "", nil
			}

			if len(result) > emacsClockMaxLen {
				result = result[:24] + "..."

			}

			return fmt.Sprintf("⌚  %s", result), nil
		},
		Span:   time.Second * 40,
		Active: emacsClockIsActive,
		Order:  44,
	})
}
