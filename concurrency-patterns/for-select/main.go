package main

import (
	"fmt"
	"time"
)

func main() {
	done := signal()

Loop:
	for {
		select {
		case <-done:
			break Loop
		default:
		}

		fmt.Println("Hello")
		time.Sleep(time.Second * 1)
	}

	fmt.Println("All done.")
}

func signal() <-chan int {
	done := make(chan int)

	go func() {
		time.Sleep(time.Second * 5)
		close(done)
	}()

	return done
}
