package plugins

import "time"

var List []Plugin

type Plugin struct {
	Span    time.Duration
	Trigger <-chan time.Time
	Getter  func() (string, error)
	Active  bool
	Cached  string
	Name    string
	Order   int
}
