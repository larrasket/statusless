package plugins

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

const emacsAwqatIsActive = true

func init() {
	PList = append(PList, Plugin{
		Getter: func() (string, error) {
			cmd := exec.Command("emacsclient", "-e", "(substring-no-properties awqat-mode-line-string)")

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

			return fmt.Sprintf(" %s", result), nil
		},
		Span:   time.Second * 40,
		Active: emacsAwqatIsActive,
		Order:  15,
	})
}
