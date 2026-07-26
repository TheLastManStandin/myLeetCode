package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a[1:3])
}

func printCodes(urls []string) {
	var wg sync.WaitGroup
	wg.Add(len(urls))

	const maxResp int = 5
	oneTimeResponces := make(chan interface{}, maxResp)

	get := func(url string, ch <-chan interface{}) {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		fmt.Println(url, resp.Status)
		defer wg.Done()

		<-ch
	}

	for _, url := range urls {
		url := url
		select {
		case oneTimeResponces <- 5:
			go get(url, oneTimeResponces)
		}
	}

	wg.Wait()
}
