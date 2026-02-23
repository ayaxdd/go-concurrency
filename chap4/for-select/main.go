package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// sendIterVarOutOnChannel()
	infLoop()
}

// inf loop waiting to be stopped
func infLoop() {
	done := make(chan any)

	timeout := func() {
		time.Sleep(1 * time.Second)
		close(done)
	}

	var wg sync.WaitGroup
	wg.Add(2)
	var firstCount, secondCount int

	// loops till done channel isn't closed
	// 1 variation
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				return
			default:
			}

			// do work
			firstCount++
		}
	}()
	// 2 variation
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				return
			default:
				// do work
				secondCount++
			}
		}
	}()

	go timeout()

	wg.Wait()
	fmt.Printf("First: %d\nSecond: %d\n", firstCount, secondCount)
}

// convert smth that can be iterated over into values on channel
func sendIterVarOutOnChannel() {
	done := make(chan any)
	chanOwner := func(done <-chan any) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for _, i := range []int{0, 1, 2, 3} {
				select {
				case <-done:
					return
				case intStream <- i:
				}
			}
		}()

		return intStream
	}

	intStream := chanOwner(done)
	for i := range intStream {
		fmt.Printf("Received: %d\n", i)
	}
}
