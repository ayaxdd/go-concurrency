package main

import (
	"fmt"
	"math/rand"
)

func example4() {
	fmt.Println("example 4:")

	done := make(chan any)
	defer close(done)

	fmt.Println("\tSimple repeat generator")
	for num := range takePipe(done, repeatGen(done, 1), 10) {
		fmt.Printf("%v ", num)
	}
	fmt.Println()

	rand := func() any {
		return rand.Int()
	}
	fmt.Println("\tRepeat generator with given func (rand)")
	for num := range takePipe(done, repeatFn(done, rand), 10) {
		fmt.Printf("%v\n", num)
	}

	var message string
	fmt.Println("\tRepeat generator with string type cast")
	for token := range toString(done, takePipe(done, repeatGen(done, "I", "am."), 5)) {
		message += token
	}
	fmt.Printf("message: %s...", message)
}

// infinite repeating numbers (untill done is called)
func repeatGen(done <-chan any, values ...any) <-chan any {
	valueStream := make(chan any)
	go func() {
		defer close(valueStream)
		for {
			for _, v := range values {
				select {
				case <-done:
				case valueStream <- v:
				}
			}
		}
	}()
	return valueStream
}

func takePipe(
	done <-chan any,
	valueStream <-chan any,
	num int,
) <-chan any {
	takeStream := make(chan any)
	go func() {
		defer close(takeStream)
		for range num {
			select {
			case <-done:
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

func repeatFn(done <-chan any, fn func() any) <-chan any {
	valueStream := make(chan any)
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			case valueStream <- fn():
			}
		}
	}()
	return valueStream
}

func toString(done <-chan any, valueStream <-chan any) <-chan string {
	stringStream := make(chan string)
	go func() {
		defer close(stringStream)
		for v := range valueStream {
			select {
			case <-done:
			case stringStream <- v.(string):
			}
		}
	}()
	return stringStream
}
