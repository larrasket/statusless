package plugins

import (
	"fmt"
	"time"

	"github.com/jezek/xgb/xproto"
)

const CapsLockIsActive = true

func init() {
	PList = append(PList, Plugin{
		Getter: func() (string, error) {
			s, err := getCapsLockState()
			if err != nil {
				return "", err
			}

			if s {
				return "CapsLck: ON", nil
			} else {
				return "", nil
			}
		},
		Span:   180 * time.Millisecond,
		Active: CapsLockIsActive,
		Order:  999,
	})
}

func getCapsLockState() (bool, error) {

	keyboard, err := xproto.GetKeyboardControl(C).Reply()
	if err != nil {
		return false, fmt.Errorf("failed to get keyboard control: %v", err)
	}
	return keyboard.LedMask&0x1 != 0, nil
}
