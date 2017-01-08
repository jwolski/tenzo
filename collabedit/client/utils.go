package main

import (
	"fmt"
	"os"
)

// Aborts program by writing message to stderr and exiting with status code 1.
func abort(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}
