package main

import (
	"fmt"
	"time"
)

func example3() {
	fmt.Println("example 3:")

	// doneFunc := func() <-chan any {
	// 	done := make(chan any)
	// 	go func() {
	// 		defer close(done)
	// 	}()
	// 	return done
	// }
	done := make(chan any)

	intStream := generatorChan(done, 1, 2, 3, 4)
	// int -> mul -> add -> mul = pipe
	pipeline := multiplyChan(done, addChan(done, multiplyChan(done, intStream, 2), 1), 2)

	start := time.Now()
	for v := range pipeline {
		fmt.Printf("result %v\n", v)
		close(done)
	}
	fmt.Println(time.Since(start))
}

func generatorChan(done <-chan any, integers ...int) <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for _, i := range integers {
			select {
			case <-done:
			case intStream <- i:
				fmt.Println(i, "generated")
			}
		}
	}()
	return intStream
}

func multiplyChan(
	done <-chan any,
	intStream <-chan int,
	multiplier int,
) <-chan int {
	multipliedStream := make(chan int)
	go func() {
		defer close(multipliedStream)
		for i := range intStream {
			select {
			case <-done:
			case multipliedStream <- i * multiplier:
				fmt.Printf("%d * %d = %d\n", i, multiplier, i*multiplier)
			}
		}
	}()
	return multipliedStream
}

func addChan(
	done <-chan any,
	intStream <-chan int,
	additive int,
) <-chan int {
	addedStream := make(chan int)
	go func() {
		defer close(addedStream)
		for i := range intStream {
			select {
			case <-done:
			case addedStream <- i + additive:
				fmt.Printf("%d + %d = %d\n", i, additive, i+additive)
			}
		}
	}()
	return addedStream
}
