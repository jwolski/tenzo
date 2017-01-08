package main

import (
	"time"
)

// Client program for collabedit application. The client awaits input from
// the user from stdin and sends the update to the document server. It also
// watches for updates from the server that have been transmitted by other
// clients.
func main() {
	// Attempt to connect to server. Abort program if connection fails.
	server := &Server{}
	if err := server.Connect(); err != nil {
		abort("Error: failed to connect to document server")
	}

	console := newConsole(server)

	// Give the server 1 second to transmit initial contents before
	// prompting the user for their input.
	update, err := server.WatchOne(1 * time.Second)
	if err == nil {
		console.SetText(update)
	}

	// Watch for document updates from document server
	watchChan, err := server.Watch()
	if err != nil {
		abort("Error: failed to set watch on document server")
	}

	// Start a background loop to read for further updates from server.
	go func() {
		for {
			update := <-watchChan
			console.SetText(update)
		}
	}()

	// This will block until the program exits.
	console.readPrintLoop()
}
