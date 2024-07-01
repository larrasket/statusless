package plugins

import (
	"log"
	"time"

	"github.com/jezek/xgb"
)

var PList []Plugin
var C *xgb.Conn

type Plugin struct {
	Span time.Duration
	// ErrorSpan is used in case if the span is too long to wait for a
	// refresh. Say for example you are getting the weather once each 2 hours,
	// however, when you started using the program there was a network issue,
	// it's better to use the ErrorSpan time instead of the Span to get the
	// weather status quicker rather than staying the two hours.
	ErrorSpan      time.Duration
	UsingErrorSpan bool
	Trigger        *time.Ticker
	Getter         func() (string, error)
	Active         bool
	Cached         string
	Name           string
	Order          int
}

func init() {
	var err error
	C, err = xgb.NewConn()
	if err != nil {
		log.Fatalf("Failed to connect to X server: %v", err)
	}
}
