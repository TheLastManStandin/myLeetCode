package main

import "fmt"

func main() {
	var myPort *int // myPort равен nil

	SetDefaultPort(myPort)

	// Что будет в myPort здесь?
	fmt.Println(myPort) // Выведет <nil>! Новое значение потерялось.
}
func SetDefaultPort(p *int) {
	if p == nil {
		p = new(int) // УБРАЛИ
		*p = 8080    // 💥 ТУТ БУДЕТ ПАНИКА!
	}
	fmt.Println("Port:", *p)
}
