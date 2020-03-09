package main

import (
	"fmt"
)

type fooBar struct {
	foo   string
	bar   string
	barCh chan bool
	fooCh chan bool
	done  chan bool
}

func (f *fooBar) outBar() {

	go func() {
		for {
			select {
			case <-f.barCh:
				fmt.Printf("%s\n", f.bar)
				select {
				case f.fooCh <- true:
				}

			case <-f.done:
				fmt.Println("breaking bar")
				break
			}
		}
	}()
}

func (f *fooBar) outFoo(sendCh, done <-chan bool) {
	go func() {
		for {
			select {
			case <-sendCh:
				fmt.Printf("%s", f.foo)
				select {
				case f.barCh <- true:
				}
				<-f.fooCh
			case <-done:
				f.done <- true
			}
		}
	}()
}

func main() {
	n := 10
	fb := fooBar{
		foo:   "foo",
		bar:   "bar",
		barCh: make(chan bool),
		fooCh: make(chan bool),
		done:  make(chan bool),
	}
	done := make(chan bool)
	sendCh := make(chan bool)
	fb.outBar()
	fb.outFoo(sendCh, done)
	forever := make(chan bool)
	go func() {
		for i := 0; i < n; i++ {
			sendCh <- true
		}

		// time.Sleep(5 * time.Second)

		select {
		case done <- true:
		}

		forever <- true
	}()
	<-forever
}
