package main

import (
	"fmt"

	"concurrency-in-go/chap4/or-channel/lib"
)

func main() {
	done := make(chan any)
	defer close(done)

	for v := range bridge(done, generate()) {
		fmt.Printf("%v ", v)
	}
}

func bridge(done <-chan any, chanStream <-chan <-chan any) <-chan any {
	valStream := make(chan any)
	go func() {
		defer close(valStream)
		for {
			var stream <-chan any
			select {
			case <-done:
				return
			case maybeStream, ok := <-chanStream:
				if !ok {
					return
				}
				stream = maybeStream
			}
			for val := range lib.OrDone(done, stream) {
				select {
				case <-done:
					return
				case valStream <- val:
				}
			}
		}
	}()
	return valStream
}

func generate() <-chan <-chan any {
	chanStream := make(chan (<-chan any))
	go func() {
		defer close(chanStream)
		for i := range 10 {
			stream := make(chan any, 1)
			stream <- i
			close(stream)
			chanStream <- stream
		}
	}()
	return chanStream
}
