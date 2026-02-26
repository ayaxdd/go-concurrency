package main

import (
	"fmt"
	"sync"
	"time"

	pipe "concurrency-in-go/chap4/or-channel/lib"
)

func main() {
	done := make(chan any)
	defer close(done)

	valueStream := generate(1, 2, 3, 4, 5, 6, 7, 8)

	chans := tee(done, valueStream, 2)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for v := range chans[0] {
			fmt.Printf("out1: %v\n", v)
		}
	}()
	go func() {
		defer wg.Done()
		for v := range chans[1] {
			time.Sleep(3 * time.Second)
			fmt.Printf("out2: %v\n", v)
		}
	}()
	wg.Wait()
}

func generate(values ...any) <-chan any {
	valueStream := make(chan any)
	go func() {
		defer close(valueStream)
		for _, v := range values {
			valueStream <- v
		}
	}()
	return valueStream
}

func tee(done, in <-chan any, numChans int) []chan any {
	chans := make([]chan any, numChans)

	for i := range numChans {
		chans[i] = make(chan any)
	}

	go func() {
		for i := range numChans {
			defer close(chans[i])
		}

		for val := range pipe.OrDone(done, in) {
			var wg sync.WaitGroup
			wg.Add(numChans)
			for i := range numChans {
				go func() {
					defer wg.Done()
					chans[i] <- val
				}()
			}
			wg.Wait()
		}
	}()

	return chans
}
