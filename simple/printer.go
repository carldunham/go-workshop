package simple

import "fmt"

// Printer just prints a message
//
type Printer string

// New makes a new SimplePrinter
//
func New(prefix string) Printer {
	return Printer(prefix)
}

// Print does the thing
//
func (simple Printer) Print(msg string) {
	fmt.Printf("Simply a message from %s: %s\n", simple, msg)
}
