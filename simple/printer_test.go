package simple

import (
	"testing"
)

func TestSimplePrinter(t *testing.T) {
	p := New("a simple printer")
	p.Print("hello, world")
}
