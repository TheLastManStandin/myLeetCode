package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// https://sobes.tech/bank/go/livecode/27ec0750-1282-41c8-8f8e-4737bf05912a

type Resp struct {
	Status int
}

func main() {
	urls := []string{"https://www.google.com",
		"https://www.facebook.com",
		"https://www.google.com", "https://www.google.com",
		"https://www.facebook.com", "https://www.google.com", "https://www.google.com", "https://www.facebook.com",
		"https://www.google.com", "https://www.google.com", "https://www.facebook.com", "https://www.google.com",
		"https://www.google.com",
	}
	printCodes(urls)
}

func printCodes(urls []string) {
	var wg sync.WaitGroup
	wg.Add(len(urls))

	const maxResp int = 1
	oneTimeResponces := make(chan interface{}, maxResp)
	defer close(oneTimeResponces)

	get := func(url string, ch <-chan interface{}) {
		defer wg.Done()
		defer func() { <-ch }()
		resp, err := http.Get(url)
		defer resp.Body.Close()
		if err != nil {
			// можно положить ошибку в канал ошибок
		}
		fmt.Println(url, resp.Status)
	}

	for _, url := range urls {
		url := url

		oneTimeResponces <- 0 // Блокирующая операция, если буффер канала заполнен
		go get(url, oneTimeResponces)
	}

	wg.Wait()
}

func Get(url string) (*Resp, error) {
	_ = url
	time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond) // Иметация работы
	return &Resp{
		Status: 3,
	}, nil
}
