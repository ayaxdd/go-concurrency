package main

import "fmt"

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
