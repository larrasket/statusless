package plugins

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

const emacsAwqatIsActive = true

func init() {
	List = append(List, Plugin{
		Getter: func() (string, error) {
			cmd := exec.Command("emacsclient", "-e", "(when (org-clocking-p) (substring-no-properties awqat-mode-line-string))")

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

			return fmt.Sprintf("%s", result), nil
		},
		Span:   time.Minute * 1,
		Active: emacsAwqatIsActive,
		Order:  15,
	})
}
