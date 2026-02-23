package main

import (
	"bytes"
	"fmt"
	"sync"
)

func lexical() {
	// exposes read-only channel
	// only channel owner can write to this channel
	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := range 6 {
				results <- i
			}
		}()

		return results
	}

	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received: %v\n", result)
		}
		fmt.Println("Done receiving")
	}

	results := chanOwner()
	consumer(results)
}

func lexicalNonChannel() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()

		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")

	// sending non intersecting subsets of data so no sync is needed
	go printData(&wg, data[:3])
	// time.Sleep(time.Second)
	go printData(&wg, data[3:])

	wg.Wait()
}
