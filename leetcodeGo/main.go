package main

import (
	"context"
	"fmt"
	"time"
)

// Имитация долгого запроса к базе данных или API
func fetchData() string {
	time.Sleep(2 * time.Second) // Долгая операция
	return "data"
}

func handleRequest() string {
	ch := make(chan string)

	go func() {
		data := fetchData()
		ch <- data // Горлышко инцидента
	}()

	// Ограничиваем время выполнения запроса одной секундой
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	select {
	case result := <-ch:
		return result
	case <-ctx.Done():
		return "timeout"
	}
}

func main() {
	// Имитируем поток запросов в веб-сервере
	for i := 0; i < 1000; i++ {
		go handleRequest()
	}
	fmt.Println("Обработка завершена")
}
