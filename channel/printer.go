package channel

import (
	"fmt"
)

// Printer prints from channels
//
type Printer struct {
	StringChannel     chan string
	EverythingChannel chan interface{}
	TriggerChannel    chan bool

	quit, done chan bool
}

// New makes a Printer
//
func New() (chp *Printer) {
	chp = &Printer{
		StringChannel:     make(chan string),
		EverythingChannel: make(chan interface{}),
		TriggerChannel:    make(chan bool),

		quit: make(chan bool, 1),
		done: make(chan bool, 1),
	}

	go func() {

		for {
			select {
			case msg := <-chp.StringChannel:
				fmt.Println(msg)

			case msg := <-chp.EverythingChannel:
				fmt.Printf("%v\n", msg)

			case <-chp.TriggerChannel:
				fmt.Println("hehe")

			case <-chp.quit:
				fmt.Println("Quitting")
				chp.done <- true
				return
			}
		}
	}()
	return
}

// Print print a complex message
func (chp *Printer) Print(msg string) {
	chp.StringChannel <- msg
}

// Quit processing messages
//
func (chp *Printer) Quit() {
	chp.quit <- true
	<-chp.done
}
