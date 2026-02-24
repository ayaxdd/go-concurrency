package main

import "testing"

func BenchmarkGeneric(b *testing.B) {
	done := make(chan any)
	defer close(done)

	b.ResetTimer()
	for range toString(done, takePipe(done, repeatGen(done, "a"), b.N)) {
	}
}

func BenchmarkTyped(b *testing.B) {
	done := make(chan any)
	defer close(done)

	b.ResetTimer()
	for range takeStr(done, repeatStr(done, "a"), b.N) {
	}
}

func takeStr(
	done <-chan any,
	valueStream <-chan string,
	num int,
) <-chan string {
	takeStream := make(chan string)
	go func() {
		defer close(takeStream)
		for i := num; i > 0 || i == -1; {
			if i != -1 {
				i--
			}
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

func repeatStr(done <-chan any, values ...string) <-chan string {
	valueStream := make(chan string)
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
