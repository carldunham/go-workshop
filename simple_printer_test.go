package main

import (
	"testing"
)

func TestSimplePrinter(t *testing.T) {
	p := SimplePrinter("a simple printer")
	p.Print("hello, world")
}
