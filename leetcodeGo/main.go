package main

import (
	"fmt"
	"sync"
)

func worker(in <-chan int, out chan<- int, workerID int) {
	for x := range in {
		fmt.Printf("worker %d, job %d\n", workerID, x)
		out <- x * x
	}
}

func main() {
	const workersNum = 3
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	results := make([]int, 0, len(data))

	// producer
	in := make(chan int)

	go func() {
		defer close(in)

		for _, x := range data {
			in <- x
		}
	}()

	// workers
	var wg sync.WaitGroup
	out := make(chan int)

	for w := 1; w <= workersNum; w++ {
		wg.Go(func() {
			worker(in, out, w)
		})
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	// consumer
	for x := range out {
		results = append(results, x)
	}

	fmt.Println("Results:", results)
}
