package main

import (
	. "git.sr.ht/~lr0/statusless/plugins"
	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
	"github.com/samber/lo"
	"log"
	"sort"
	"strings"
	"sync"
	"time"
)

const sep = "  ┃  "
const plain = "┃"

// TODO priorities. I think each plugin should have its priority.
func main() {

	PList = lo.Filter(PList, func(item Plugin, index int) bool {
		return item.Active
	})

	sort.Slice(PList, func(i, j int) bool {
		return PList[i].Order > PList[j].Order
	})

	// init the latest value for all plugins and the ticker
	for i, p := range PList {

		s, err := p.Getter()
		if err != nil {
			log.Printf("%s: %s\n", p.Name, err.Error())
			if PList[i].ErrorSpan != (0 * time.Second) {
				PList[i].Trigger = time.NewTicker(p.ErrorSpan)
				PList[i].UsingErrorSpan = true
			}
		}
		PList[i].Cached = s
		if !PList[i].UsingErrorSpan {
			PList[i].Trigger = time.NewTicker(p.Span)
		}
	}

	// set first bar
	updateXroot(C, makeBar())

	var wg sync.WaitGroup
	for i := range PList {
		go func(p *Plugin) {
			for {
				select {
				case <-p.Trigger.C:
					s, err := p.Getter()
					if err != nil {
						s = err.Error()
						log.Printf("%s: %s\n", p.Name, s)
						if p.ErrorSpan != (0 * time.Second) {
							p.UsingErrorSpan = true
							p.Trigger.Stop()
							p.Trigger = time.NewTicker(p.ErrorSpan)
						}
					} else if p.UsingErrorSpan {
						p.UsingErrorSpan = false
						p.Trigger.Stop()
						p.Trigger = time.NewTicker(p.Span)
					}
					p.Cached = s
					updateXroot(C, makeBar())
				}
			}
		}(&PList[i])
	}
	wg.Add(1)
	wg.Wait()

}

func makeBar() string {
	var s strings.Builder
	l := len(PList) - 1
	s.WriteString("  ")
	for i, p := range PList {
		if p.Cached == "" {
			continue
		}

		s.WriteString(p.Cached)
		if i == l {
			s.WriteString(" " + plain)
		} else {
			s.WriteString(sep)
		}

	}
	return s.String()
}

func updateXroot(conn *xgb.Conn, s string) {
	err := xproto.ChangePropertyChecked(
		conn,
		xproto.PropModeReplace,
		xproto.Setup(conn).DefaultScreen(conn).Root,
		xproto.AtomWmName,
		xproto.AtomString,
		8, // Format (8-bit)
		uint32(len(s)),
		[]byte(s),
	).Check()
	if err != nil {
		log.Fatalf("Failed to change root window name: %v", err)
	}
}
