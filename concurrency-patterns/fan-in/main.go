package main

import "sync"

func fanIn(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	unifiedStream := make(chan interface{})

	unify := func(c <-chan interface{}) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case unifiedStream <- i:
			}
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go unify(c)
	}

	go func() {
		wg.Wait()
		close(unifiedStream)
	}()

	return unifiedStream
}
