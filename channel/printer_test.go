package channel

import (
	"testing"
)

func TestPrinter(t *testing.T) {
	c := New()
	defer func() {
		c.Quit()
	}()

	c.Print("hello, world")

	c.TriggerChannel <- true

	c.EverythingChannel <- 42
	c.EverythingChannel <- true
	c.EverythingChannel <- []string{"hi", "there"}

	c.TriggerChannel <- true
}
