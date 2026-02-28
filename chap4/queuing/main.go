package main

import (
	"fmt"
	"time"
)

func main() {
	run()
}

func run() {
	done := make(chan any)
	defer close(done)

	start := time.Now()

	go func() {
		for {
			fmt.Println(time.Since(start))
			time.Sleep(time.Second)
		}
	}()

	zeros := take(done, repeat(done, 0), 3)
	short := sleep(done, zeros, 1*time.Second)
	long := sleep(done, short, 4*time.Second)
	pipe := long

	for p := range pipe {
		fmt.Println(p)
	}
	fmt.Printf("time pass: %v", time.Since(start))
}

func repeat(done <-chan any, value any) <-chan any {
	repeatStream := make(chan any)
	go func() {
		defer close(repeatStream)
		for {
			select {
			case <-done:
				return
			case repeatStream <- value:
			}
		}
	}()
	return repeatStream
}

func take(done, valueStream <-chan any, numTake int) <-chan any {
	takeStream := make(chan any)
	go func() {
		defer close(takeStream)
		for range numTake {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

func sleep(done, valueStream <-chan any, delay time.Duration) <-chan any {
	delayStream := make(chan any)
	go func() {
		defer close(delayStream)
		for val := range valueStream {
			time.Sleep(delay)
			select {
			case <-done:
				return
			case delayStream <- val:
			}
		}
	}()
	return delayStream
}
