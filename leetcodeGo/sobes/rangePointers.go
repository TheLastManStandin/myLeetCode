package main

import (
	"fmt"
)

func main() {
	var ptrs []*int // срез указателей
	for _, v := range []int{10, 20, 30, 40} {
		ptrs = append(ptrs, &v) // сохраняем адрес переменной‑итератора
	}

	for _, p := range ptrs {
		fmt.Printf("%d ", *p) // выводим значение, на которое указывает каждый указатель
	}
}
