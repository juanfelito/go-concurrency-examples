package main

import "fmt"

func main() {
	done := make(chan interface{})
	defer close(done)

	values := []int{1, 2, 3, 4, 5}

	inputStream := generator(done, values)

	result := multiply(done, inputStream, 2)

	for v := range result {
		fmt.Println(v)
	}
}

// Function to turn a slice into a stream
func generator(done <-chan interface{}, elements []int) <-chan int {
	stream := make(chan int)

	go func() {
		defer close(stream)
		for _, element := range elements {
			select {
			case <-done:
				return
			case stream <- element:
			}
		}
	}()

	return stream
}

func multiply(done <-chan interface{}, stream <-chan int, multiplier int) <-chan interface{} {
	resultStream := make(chan interface{})
	go func() {
		defer close(resultStream)
		for value := range orDone(done, stream) {
			intValue := value.(int)
			resultStream <- intValue * multiplier
		}
	}()

	return resultStream
}
