package plugins

import "time"

var List []plugin

type plugin struct {
	Trigger time.Duration
	Getter  func() (string, error)
	Active  bool
	Cached  *string
	Name    string
}
