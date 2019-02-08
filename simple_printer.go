package main

import "fmt"

// SimplePrinter just prints a message
//
type SimplePrinter string

// Print does the thing
//
func (simple SimplePrinter) Print(msg string) {
	fmt.Printf("Simply a message from %s: %s\n", simple, msg)
}
