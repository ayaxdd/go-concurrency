package main

import (
	"fmt"
	"time"
)

func main() {
	// simpleSelect()
	// silmultaneousReading()
	forSelect()
}

func emptySelect() {
	select {} // block forever
}

func forSelect() {
	done := make(chan any)
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	workCounter := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}

		// simulate work
		workCounter++
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("Achieved %d cycles of work before signalled to stop\n", workCounter)
}

func silmultaneousReading() {
	c1 := make(chan any)
	c2 := make(chan any)
	close(c1)
	close(c2)

	var c1Count, c2Count int
	for range 1000000 {
		// both closed => select pseudorandom
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}
	}

	fmt.Printf("c1Count: %d\nc2Count: %d\n", c1Count, c2Count)
}

func simpleSelect() {
	start := time.Now()
	c := make(chan any)
	go func() {
		time.Sleep(5 * time.Second)
		close(c)
	}()

	fmt.Println("Blocking on read...")
	select {
	case <-c:
		fmt.Printf("Unblocked %v later\n", time.Since(start))
	case <-time.After(time.Second):
		fmt.Println("Timeout")
	}
}
