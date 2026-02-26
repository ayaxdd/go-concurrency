package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"

	pipe "concurrency-in-go/chap4/pipelines/lib"
)

func main() {
	rand := func() any { return rand.Intn(50000000) }
	// example1(rand)
	example2(rand)
}

func example1(fn func() any) {
	fmt.Println("example 1:")

	done := make(chan any)
	defer close(done)

	start := time.Now()

	randIntStream := pipe.ToInt(done, pipe.RepeatFn(done, fn))
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

func example2(fn func() any) {
	fmt.Println("example 2:")

	done := make(chan any)
	defer close(done)

	start := time.Now()

	randIntStream := pipe.ToInt(done, pipe.RepeatFn(done, fn))
	// fan-out
	numFinders := runtime.NumCPU()
	fmt.Printf("Spinning up %d prime finders\n", numFinders)
	finders := make([]<-chan any, numFinders)
	fmt.Println("Primes:")
	for i := range numFinders {
		finders[i] = primeFinder(done, randIntStream)
	}

	for prime := range pipe.Take(done, fanIn(done, finders...), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v", time.Since(start))
}

func fanIn(done <-chan any, channels ...<-chan any) <-chan any {
	var wg sync.WaitGroup
	muxStream := make(chan any)

	mux := func(c <-chan any) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case muxStream <- i:
			}
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go mux(c)
	}

	go func() {
		wg.Wait()
		close(muxStream)
	}()

	return muxStream
}
