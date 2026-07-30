package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	alreadyStored := make(map[int]struct{})
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	capacity := 1000
	doubles := make([]int, 0, capacity)
	for i := 0; i < capacity; i++ {
		doubles = append(doubles, rand.Intn(10)) // create rand num 0..9
	}
	// 3, 4, 5, 0, 4, 9, 9, 8, 6, 6, 5, 5, 4, 4, 2, 1, 2, 3, 1 ...

	uniqueIDs := make(chan int, capacity)

	for i := 0; i < capacity; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			_, ok := alreadyStored[doubles[i]]
			if !ok {
				alreadyStored[doubles[i]] = struct{}{}
			}
			mu.Unlock()
			if !ok {
				uniqueIDs <- doubles[i]
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(uniqueIDs)
	}()

	for val := range uniqueIDs {
		fmt.Println(val)
	}
}

// BAD \ OLD SOLUTION

//func main() {
//	//var alreadyStored map[int]struct{}
//	alreadyStored := make(map[int]struct{})
//	mu := sync.Mutex{}
//	capacity := 1000
//	counter := 10
//
//	doubles := make([]int, 0, capacity)
//	for i := 0; i < capacity; i++ {
//		doubles = append(doubles, rand.Intn(10)) // create rand num 0..9
//	}
//	// 3, 4, 5, 0, 4, 9, 9, 8, 6, 6, 5, 5, 4, 4, 2, 1, 2, 3, 1 ...
//
//	uniqueIDs := make(chan int, capacity)
//	wg := sync.WaitGroup{}
//
//	go func() {
//		for val := range uniqueIDs {
//			fmt.Println(val)
//		}
//	}()
//
//	for i := 0; i < capacity; i++ {
//		wg.Add(1)
//		go func(i int) {
//			defer wg.Done()
//			mu.Lock()
//			if _, ok := alreadyStored[doubles[i]]; !ok {
//				alreadyStored[doubles[i]] = struct{}{}
//				counter--
//				uniqueIDs <- doubles[i]
//			}
//			mu.Unlock()
//
//			if counter == 0 {
//				close(uniqueIDs)
//			}
//		}(i)
//	}
//
//	wg.Wait()
//	//fmt.Println(uniqueIDs)
//}
