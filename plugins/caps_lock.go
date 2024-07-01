package plugins

import (
	"fmt"
	"time"

	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
)

const CapsLockIsActive = true

func init() {
	PList = append(PList, Plugin{
		Getter: func() (string, error) {
			s, err := capsLock(C)
			if err != nil {
				return "", err
			}

			if s == 2 {
				return "", nil
			} else if s == 3 {
				return "CapsLck: ON", nil
			}
			return fmt.Sprintf("UNKOWN Signal %d", s), nil
		},
		Span:   180 * time.Millisecond,
		Active: CapsLockIsActive,
		Order:  999,
	})
}

func capsLock(conn *xgb.Conn) (uint32, error) {
	state, err := xproto.GetKeyboardControl(conn).Reply()
	if err != nil {
		return 0, err
	}
	return state.LedMask, err
}
