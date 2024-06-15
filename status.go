package main

import (
	"fmt"
	"log"

	"git.sr.ht/~lr0/status.go/plugins"
	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
)

func main() {
	// Connect to the X server
	conn, err := xgb.NewConn()
	if err != nil {
		log.Fatalf("Failed to connect to X server: %v", err)
	}
	defer conn.Close()

	// Get the root window
	setup := xproto.Setup(conn)
	root := setup.DefaultScreen(conn).Root

	// New name for the root window
	newName := "New Root Window Name"

	// Change the property

	for _, value := range plugins.List {
		t, err := value.Getter()
		if err != nil {
			panic(err)
		}

		fmt.Println(t)
	}

	err = xproto.ChangePropertyChecked(
		conn,
		xproto.PropModeReplace,
		root,
		xproto.AtomWmName,
		xproto.AtomString,
		8, // Format (8-bit)
		uint32(len(newName)),
		[]byte(newName),
	).Check()
	if err != nil {
		log.Fatalf("Failed to change root window name: %v", err)
	}

}
