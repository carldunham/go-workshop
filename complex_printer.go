package main

import "fmt"

// ComplexPrinter is a complex printer
//
type ComplexPrinter struct {
	Prefix string
	Suffix string

	counter int
}

// Print print a complex message
func (complex *ComplexPrinter) Print(msg string) {
	fmt.Printf("%d: %s %s %s\n", complex.counter, complex.Prefix, msg, complex.Suffix)

	complex.counter++
}

// NewComplexPrinter makes a ComplexPrinter
//
func NewComplexPrinter(prefix, suffix string) *ComplexPrinter {
	return &ComplexPrinter{
		Prefix: prefix,
		Suffix: suffix,
	}
}

// Count returns the number of time this ComplexPrinter was called
//
func (complex *ComplexPrinter) Count() int {
	return complex.counter
}
