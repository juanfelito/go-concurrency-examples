package main

func orDone(done <-chan interface{}, c <-chan int) <-chan interface{} {
	stream := make(chan interface{})

	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if ok == false {
					return
				}
				select {
				case stream <- v:
				case <-done:
				}
			}
		}
	}()

	return stream
}
