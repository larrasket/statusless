package plugins

import (
	"fmt"
	"time"
)

// '%d %b, %H:%M:%S %p'
const timeFormat = "02 Jan, 15:04:05 PM"
const dateIsActive = true

func init() {
	List = append(List, plugin{
		Getter: func() (string, error) {

			return fmt.Sprintf(" %s", time.Now().Format(timeFormat)), nil
		},
		Trigger: time.Second,
		Active:  dateIsActive,
	})
}
