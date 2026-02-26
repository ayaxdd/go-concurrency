package main

import (
	"fmt"
	"math/rand"
	"time"

	pipe "concurrency-in-go/chap4/pipelines/lib"
)

func main() {
	rand := func() any { return rand.Intn(50) }

	done := make(chan any)
	defer close(done)

	start := time.Now()

	randIntStream := pipe.ToInt(done, pipe.RepeatFn(done, rand))
	fmt.Println("Primes:")
	for prime := range pipe.Take(done, primeFinder(done, randIntStream), 10) {
		fmt.Printf("\t%d\n", prime)
	}
	fmt.Printf("Search took: %v", time.Since(start))
}

// inefficient
func primeFinder(done <-chan any, intStream <-chan int) <-chan any {
	primeStream := make(chan any)
	go func() {
		defer close(primeStream)
		for i := range intStream {
			ok := isPrime(i)
			if !ok {
				continue
			}

			select {
			case <-done:
				return
			case primeStream <- i:
			}
		}
	}()
	return primeStream
}

func isPrime(i int) bool {
	if i < 2 {
		return false
	}
	for k := i - 1; k > 1; k-- {
		if i%k == 0 {
			return false
		}
	}
	return true
}
