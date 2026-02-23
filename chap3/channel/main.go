package main

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	channelOwnership()
}

// keep channel scope as small as possible
func channelOwnership() {
	// encapsulate channel
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5)
		go func() {
			// owning goroutine must close and write in channel
			defer close(resultStream)
			for i := range 6 {
				resultStream <- i
			}
			time.Sleep(time.Second * 5)
		}()

		return resultStream
	}

	resultStream := chanOwner()
	// blocking: read till channel is closed
	for result := range resultStream {
		fmt.Printf("Recieved: %d\n", result)
	}
	fmt.Println("Done")
}

func nilChannel() {
	var dataStream chan any
	<-dataStream             // deadlock
	dataStream <- struct{}{} // deadlock
	close(dataStream)        // panic
}

func bufferedChannel() {
	var stdoutBuf bytes.Buffer
	defer stdoutBuf.WriteTo(os.Stdout)

	intStream := make(chan int, 3)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuf, "Producer Done")
		for i := range 5 {
			fmt.Fprintf(&stdoutBuf, "Sending: %d\n", i+1)
			intStream <- i + 1
		}
	}()

	for i := range intStream {
		fmt.Fprintf(&stdoutBuf, "Recieved: %d\n", i)
	}
}

func unblockingNGoroutines() {
	begin := make(chan any)
	var wg sync.WaitGroup
	for i := range 5 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines...")
	close(begin)
	wg.Wait()
}

func rangeOverChannel() {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := range 5 {
			intStream <- i + 1
		}
	}()

	go func() {
		for i := 10; i < 20; i++ {
			intStream <- i
		}
	}()

	// break the loop when channel is closed
	for i := range intStream {
		fmt.Println(i)
	}
}

func closedChanValue() {
	stringStream := make(chan string)
	go func() {
		// stringStream <- "Hello"
		close(stringStream)
	}()
	salutation, ok := <-stringStream
	fmt.Printf("(%v): %v\n", ok, salutation)
}

func deadlock() {
	stringStream := make(chan string)
	a := func() {
		if 0 != 1 {
			return
		}
		stringStream <- "hello!"
	}
	b := func() {
		fmt.Println(<-stringStream)
	}

	go a()
	go b()

	time.Sleep(3 * time.Second)
}
