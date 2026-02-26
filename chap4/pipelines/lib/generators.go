package lib

// infinite repeating numbers (untill done is called)
func Repeat(done <-chan any, values ...any) <-chan any {
	valueStream := make(chan any)
	go func() {
		defer close(valueStream)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case valueStream <- v:
				}
			}
		}
	}()
	return valueStream
}

func Take(
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
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

func RepeatFn(done <-chan any, fn func() any) <-chan any {
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

func ToString(done <-chan any, valueStream <-chan any) <-chan string {
	stringStream := make(chan string)
	go func() {
		defer close(stringStream)
		for v := range valueStream {
			select {
			case <-done:
				return
			case stringStream <- v.(string):
			}
		}
	}()
	return stringStream
}

func ToInt(done <-chan any, valueStream <-chan any) <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for v := range valueStream {
			select {
			case <-done:
				return
			case intStream <- v.(int):
			}
		}
	}()
	return intStream
}
