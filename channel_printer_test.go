package main

import (
	"testing"
)

func TestChannelPrinter(t *testing.T) {
	c := NewChannelPrinter()
	defer c.Quit()

	c.Print("hello, world")

	c.TriggerChannel <- true

	c.EverythingChannel <- 42
	c.EverythingChannel <- true
	c.EverythingChannel <- []string{"hi", "there"}

	c.TriggerChannel <- true
}
