package main

import (
	"fmt"
	"math/rand"
	"time"

	gen "concurrency-in-go/chap4/pipelines/lib"
)

func main() {
	example1()
	example2()
	example3()
	example4()
}

func example1() {
	fmt.Println("example 1:")

	ints := []int{1, 2, 3, 4}
	// for _, v := range addBatch(multiplyBatch(ints, 2), 1) { // v*2+1
	// 	fmt.Printf("%v\n", v)
	// }
	for _, v := range multiplyBatch(addBatch(multiplyBatch(ints, 2), 1), 2) { // (v*2+1)*2
		fmt.Printf("%v\n", v)
	}
}

// stage 1
func multiplyBatch(values []int, multiplier int) []int {
	multipliedValues := make([]int, len(values))
	for i, v := range values {
		multipliedValues[i] = v * multiplier
	}
	return multipliedValues
}

// stage 2
func addBatch(values []int, additive int) []int {
	addedValues := make([]int, len(values))
	for i, v := range values {
		addedValues[i] = v + additive
	}
	return addedValues
}

func example2() {
	fmt.Println("example 2:")

	ints := []int{1, 2, 3, 4}
	for _, v := range ints {
		fmt.Printf("%v\n", multiplyStream(addStream(multiplyStream(v, 2), 1), 2)) // (v*2+1)*2
	}
}

func multiplyStream(value, multiplier int) int {
	return value * multiplier
}

func addStream(value, additive int) int {
	return value + additive
}

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

func example4() {
	fmt.Println("example 4:")

	done := make(chan any)
	defer close(done)

	fmt.Println("\tSimple repeat generator")
	for num := range gen.Take(done, gen.Repeat(done, 1), 10) {
		fmt.Printf("%v ", num)
	}
	fmt.Println()

	rand := func() any {
		return rand.Int()
	}
	fmt.Println("\tRepeat generator with given func (rand)")
	for num := range gen.Take(done, gen.RepeatFn(done, rand), 10) {
		fmt.Printf("%v\n", num)
	}

	var message string
	fmt.Println("\tRepeat generator with string type cast")
	for token := range gen.ToString(done, gen.Take(done, gen.Repeat(done, "I", "am."), 5)) {
		message += token
	}
	fmt.Printf("message: %s...", message)
}
