package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()
	
	for {
		select {
		case v, ok := <-ch:
			fmt.Println(v)
		case <-time.After(3 * time.Second):
			fmt.Println("timeout")
			break
		}
	}
}

/*
1
2
3
4
5
timeout
*/
