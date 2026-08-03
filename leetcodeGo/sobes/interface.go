package main

import "fmt"

type Runner interface {
	Run()
}

type Swimmer interface {
	Swim()
}

type Human struct{}

func (h Human) Run() {}

type Baby struct{}

func (b Baby) Swim() {}

func main() {
	vasya := Human{}
	var r Runner
	r = vasya

	// Шаг 1
	s1, ok1 := r.(Swimmer)
	fmt.Printf("s1: %v, ok1: %v\n", s1, ok1)

	// Шаг 2
	//s2, ok2 := r.(Baby)
	//fmt.Printf("s2: %v, ok2: %v\n", s2, ok2)

	// Шаг 3
	//var i interface{} = r
	//s3, ok3 := i.(Swimmer)
	//fmt.Printf("s3: %v, ok3: %v\n", s3, ok3)

	// Шаг 4
	//s4, ok4 := i.(Baby)
	//fmt.Printf("s4: %v, ok4: %v\n", s4, ok4)
}
