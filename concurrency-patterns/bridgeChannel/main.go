package main

func Bridge(done <-chan interface{}, chanStream <-chan <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})

	go func() {
		defer close(valStream)
		for {
			var stream <-chan interface{}
			select {
			case maybeChan, ok := <-chanStream:
				if ok == false {
					return
				}
				stream = maybeChan
			case <-done:
				return
			}

			for val := range orDone(done, stream) {
				select {
				case valStream <- val:
				case <-done:
				}
			}
		}
	}()

	return valStream
}

func orDone(done <-chan interface{}, c <-chan interface{}) <-chan interface{} {
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
