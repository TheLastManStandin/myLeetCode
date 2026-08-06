package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)
	return rnd
}

func predictableFunc() (int64, error) {
	t := time.NewTimer(3 * time.Second) // в старых версиях нужно было следить за утечкой
	ch := make(chan int64, 1)
	funcStart := time.Time{}

	go func() {
		funcStart = time.Now()
		ch <- unpredictableFunc()
		close(ch)
	}()

	select {
	case <-t.C:
		return 0, errors.New("timed out")
	case val := <-ch:
		fmt.Println(time.Since(funcStart))
		return val, nil
	}
}

func main() {
	fmt.Println("Started")
	fmt.Println(predictableFunc())
}
