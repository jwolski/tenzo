package main

import (
	"bufio"
	"fmt"
	"os"
)

type Console struct {
	server       *Server
	inputPrompt  string
	serverString string
	atPrompt     bool
}

func newConsole(server *Server) *Console {
	serverString := fmt.Sprintf("%s>", server.conn.RemoteAddr().String())
	return &Console{
		server:       server,
		serverString: serverString,
		inputPrompt:  "$ ",
		atPrompt:     false,
	}
}

func (c *Console) printPrompt() {
	// Bypass printing the prompt multiple times.
	if c.atPrompt {
		return
	}
	fmt.Print(c.inputPrompt)
	c.atPrompt = true
}

func (c *Console) readPrintLoop() {
	// Read user input from stdin and send to server forever and ever until
	// program exits
	for {
		c.printPrompt()

		// Read user input from stdin. Abort the program if there is any error.
		stdinReader := bufio.NewReader(os.Stdin)
		input, err := stdinReader.ReadString([]byte("\n")[0])
		if err != nil {
			abort("Error: failed to read from stdin")
		}

		// Write input back to server
		err = c.server.Send(input)
		if err != nil {
			abort("Error: failed to update server")
		}

		// Reset prompt for next iteration
		c.atPrompt = false
	}
}

func (c *Console) SetText(text string) {
	// If we're prompting the user, print the text on a line underneath the
	// prompt.
	if c.atPrompt {
		fmt.Println()
	}

	// Print the text and prompt the user again.
	fmt.Println(c.serverString, text)
	c.atPrompt = false
	c.printPrompt()
}
