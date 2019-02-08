package complex

import (
	"fmt"
	"testing"
)

func TestComplexPrinter(t *testing.T) {
	c := New("pre", "post")

	if c.Count() != 0 {
		t.Errorf("Expected initial count of 0, got %d instead", c.Count())
	}

	for i := 0; i < 10; i++ {
		c.Print(fmt.Sprintf("hello, world #%d", i))

		if c.Count() != i+1 {
			t.Errorf("Expected count of %d, got %d instead", i+1, c.Count())
		}
	}
}
