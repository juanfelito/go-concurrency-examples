package main

import (
	"bytes"
	"fmt"
	"sync"
)

func main() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()

		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Print(buff.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("sampletext")

	go printData(&wg, data[5:])
	go printData(&wg, data[:5])

	wg.Wait()
}
