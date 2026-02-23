package main

import (
	"fmt"
	"sync"
)

func main() {
	var calcsCreated int
	calcPool := &sync.Pool{
		New: func() interface{} {
			calcsCreated++
			mem := make([]byte, 1024)
			return &mem
		},
	}

	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const workers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(workers)

	for i := workers; i > 0; i-- {
		go func() {
			defer wg.Done()
			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)
		}()
	}

	wg.Wait()
	fmt.Printf("%d calcs were created", calcsCreated)
}

