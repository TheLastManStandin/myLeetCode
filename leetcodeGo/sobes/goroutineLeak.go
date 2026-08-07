package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		requestData(1)
	}

	time.Sleep(time.Second * 1)
	fmt.Printf("Number of hanging goroutines: %d\n", runtime.NumGoroutine())
}

func requestData(timeout time.Duration) string {
	dataChan := make(chan string, 1)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	go func() {
		dataChan <- requestFromSlowServer(ctx)
	}()

	select {
	case result := <-dataChan:
		fmt.Printf("[+] request returned: %s\n", result)
		return result
	case <-ctx.Done():
		fmt.Println("[!!] request timeout!")
		return ""
	}
}

func requestFromSlowServer(ctx context.Context) string {
	select {
	case <-ctx.Done():
		return ""
	case <-time.After(time.Second * 1):
		return "very important data"
	}
}
