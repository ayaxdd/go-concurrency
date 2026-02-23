package main

import (
	"fmt"
	"sync"
)

type button struct {
	clicked *sync.Cond
}

func main() {
	btn := button{clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)
	subscribe(btn.clicked, func() {
		fmt.Println("Maximizing window")
		clickRegistered.Done()
	})
	subscribe(btn.clicked, func() {
		fmt.Println("Displaying annoying dialog box")
		clickRegistered.Done()
	})
	subscribe(btn.clicked, func() {
		fmt.Println("Mouse clicked")
		clickRegistered.Done()
	})

	btn.clicked.Broadcast()

	clickRegistered.Wait()
}
