package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]int)

	mu := sync.Mutex{}

	wg := sync.WaitGroup{}
	res := make(chan int)

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
			res <- i
		}(i)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	for v := range res {
		fmt.Println(v)
	}
}
