package main

import (
	"log"
	"sort"
	"strings"
	"sync"
	"time"

	"git.sr.ht/~lr0/statusless/plugins"
	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
)

const sep = "  ┃  "
const plain = "┃"

func main() {
	sort.Slice(plugins.List, func(i, j int) bool {
		return plugins.List[i].Order > plugins.List[j].Order
	})

	// init the latest value for all plugins and the ticker
	for i, p := range plugins.List {
		s, err := p.Getter()
		if err != nil {
			log.Fatalf("%s: %s\n", p.Name, err.Error())
		}
		plugins.List[i].Cached = s
		plugins.List[i].Trigger = time.NewTicker(p.Span)
	}

	// set first bar
	updateXroot(makeBar())

	var wg sync.WaitGroup
	for i := range plugins.List {
		go func(p *plugins.Plugin) {
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
					updateXroot(makeBar())
				}
			}
		}(&plugins.List[i])
	}
	wg.Add(1)
	wg.Wait()

}

func makeBar() string {
	var s strings.Builder
	l := len(plugins.List) - 1
	s.WriteString("  ")
	for i, p := range plugins.List {
		if !p.Active || p.Cached == "" {
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

func updateXroot(s string) {
	conn, err := xgb.NewConn()
	if err != nil {
		log.Fatalf("Failed to connect to X server: %v", err)
	}
	defer conn.Close()

	setup := xproto.Setup(conn)
	root := setup.DefaultScreen(conn).Root

	err = xproto.ChangePropertyChecked(
		conn,
		xproto.PropModeReplace,
		root,
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
