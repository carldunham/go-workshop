package main

import (
	"workshop/channel"
	"workshop/complex"
	"workshop/simple"
)

func main() {
	var ptr Printer

	ptr = simple.New("here")
	ptr.Print("hello, world")

	ptr = complex.NewComplexPrinter("pre", "post")
	ptr.Print("hello, world")

	ptr = channel.New()
	ptr.Print("hello, world")
}
