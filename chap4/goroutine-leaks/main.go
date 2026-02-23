package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// exampleLeak()
	// preventingLeak()
	// writeBlockLeak()
	preventingWriteBlockLeak()
}

func preventingWriteBlockLeak() {
	newRandStream := func(done <-chan any) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited")
			defer close(randStream)
			for {
				select {
				case <-done:
					return
				case randStream <- rand.Int():
					// do smth
				}
			}
		}()
		return randStream
	}

	done := make(chan any)
	randStream := newRandStream(done)

	go func() {
		defer close(done)
		fmt.Println("3 random ints:")
		for i := range 3 {
			fmt.Printf("%d: %d\n", i, <-randStream)
		}
	}()

	time.Sleep(1 * time.Second)
}

func writeBlockLeak() {
	newRandStream := func() <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited")
			defer close(randStream)
			for {
				randStream <- rand.Int()
			}
		}()
		return randStream
	}

	randStream := newRandStream()
	fmt.Println("3 random ints:")
	for i := range 3 {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
}

func preventingLeak() {
	doWork := func(done <-chan any, strings <-chan string) <-chan any {
		terminated := make(chan any)
		go func() {
			defer fmt.Println("doWork exited")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					_ = s
					// do smth with s
				case <-done:
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan any)
	terminated := doWork(done, nil)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()

	<-terminated // wait till doWork exit
	fmt.Println("Done")
}

func exampleLeak() {
	doWork := func(strings <-chan string) <-chan any {
		completed := make(chan any)
		go func() {
			defer fmt.Println("doWork exited")
			defer close(completed)
			for s := range strings {
				// do some work
				fmt.Println(s)
			}
		}()
		return completed
	}

	comp := doWork(nil)
	_ = comp
	// more work
	// <-comp // deadlock
	fmt.Println("Done")
}
