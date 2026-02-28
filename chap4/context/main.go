package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// doneCtx()
	stdCtx()
	ProcessRequest(context.Background(), "jane", "abc12")
}

func doneCtx() {
	var wg sync.WaitGroup
	done := make(chan any)
	defer close(done)

	wg.Go(func() {
		if err := printGreeting(done); err != nil {
			fmt.Printf("%v", err)
			return
		}
	})

	wg.Go(func() {
		if err := printFarewell(done); err != nil {
			fmt.Printf("%v", err)
			return
		}
	})

	wg.Wait()
}

func stdCtx() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Go(func() {
		if err := printGreetingCtx(ctx); err != nil {
			fmt.Printf("cannot print greeting: %v\n", err)
			cancel()
		}
	})

	wg.Go(func() {
		if err := printFarewellCtx(ctx); err != nil {
			fmt.Printf("cannot print farewell: %v\n", err)
			cancel()
		}
	})

	wg.Wait()
}

func printGreeting(done <-chan any) error {
	greeting, err := genGreeting(done)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", greeting)
	return nil
}

func printGreetingCtx(ctx context.Context) error {
	greeting, err := genGreetingCtx(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", greeting)
	return nil
}

func printFarewell(done <-chan any) error {
	farewell, err := genFarewell(done)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", farewell)
	return nil
}

func printFarewellCtx(ctx context.Context) error {
	farewell, err := genFarewellCtx(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", farewell)
	return nil
}

func genGreeting(done <-chan any) (string, error) {
	switch locale, err := locale(done); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "hello", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func genGreetingCtx(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	switch locale, err := localeCtx(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "hello", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func genFarewell(done <-chan any) (string, error) {
	switch locale, err := locale(done); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "goodbye", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func genFarewellCtx(ctx context.Context) (string, error) {
	switch locale, err := localeCtx(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "goodbye", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func locale(done <-chan any) (string, error) {
	select {
	case <-done:
		return "", fmt.Errorf("canceled")
	case <-time.After(5 * time.Second):
	}
	return "EN/US", nil
}

func localeCtx(ctx context.Context) (string, error) {
	// fail fast
	if deadline, ok := ctx.Deadline(); ok {
		if deadline.Sub(time.Now().Add(5*time.Second)) <= 0 {
			return "", context.DeadlineExceeded
		}
	}

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(5 * time.Second):
	}
	return "EN/US", nil
}
