package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	requestsStatuses([]string{"https://google.com", "https://instagram.com"})
}

func requestsStatuses(urls []string) {
	wg := sync.WaitGroup{}
	wg.Add(len(urls))

	for _, url := range urls {
		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer resp.Body.Close()

			fmt.Println("URL:", url, "StatusCode:", resp.StatusCode)
		}(url)
	}

	wg.Wait()
	fmt.Println("finish")
}

// BAD \ OLD SOLUTION

//func paralel(urls []string) {
//	wg := sync.WaitGroup{}
//	wg.Add(len(urls))
//	for i := 0; i < len(urls); i++ {
//		go func() {
//			defer wg.Done()
//			fmt.Println(urls[i])
//			resp, err := http.Get(urls[i])
//			if err != nil {
//				fmt.Println(err)
//				return
//			}
//			fmt.Println(urls[i], resp.StatusCode)
//		}()
//	}
//	wg.Wait()
//	fmt.Println("done")
//}
