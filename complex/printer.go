package complex

import "fmt"

// Printer is a complex printer
//
type Printer struct {
	Prefix string
	Suffix string

	counter int
}

// New makes a Printer
//
func New(prefix, suffix string) *Printer {
	return &Printer{
		Prefix: prefix,
		Suffix: suffix,
	}
}

// Print print a complex message
func (complex *Printer) Print(msg string) {
	fmt.Printf("%d: %s %s %s\n", complex.counter, complex.Prefix, msg, complex.Suffix)

	complex.counter++
}

// Count returns the number of time this Printer was called
//
func (complex *Printer) Count() int {
	return complex.counter
}
