package main

import "fmt"

func main() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})

		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)

			for s := range strings {
				fmt.Println(s)
			}
		}()

		return completed
	}

	// passing nil will cause doWork to never close it's channel, making it live the whole lifespan of the program.
	doWork(nil)

	fmt.Println("Done.")
}
