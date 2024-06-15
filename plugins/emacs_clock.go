package plugins

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

const emacsClockIsActive = true

func init() {
	List = append(List, Plugin{
		Getter: func() (string, error) {
			cmd := exec.Command("emacsclient", "-e", "(substring-no-properties (org-clock-get-clock-string))")

			output, err := cmd.Output()
			if err != nil {
				return "", err
			}

			result := strings.TrimSpace(string(output))
			result = strings.Trim(result, "\"")
			result = strings.NewReplacer("(", "", ")", "").Replace(result)

			return fmt.Sprintf("⌚  %s", result), nil
		},
		Span:   time.Minute * 2,
		Active: emacsClockIsActive,
		Order:  44,
	})
}
