package main

import (
	"fmt"
	"unsafe"
)

// https://sobes.tech/bank/go/livecode/4899b920-211b-4015-aa38-ddd55f573c7e

type st struct {
	// 4 ??? sizeof
	//p1 bool
	//p2 int //64
	//p3 bool
}

// ptr
func main() {
	myStr := st{}
	//
	//fmt.Println(myStr) // 1 ???
	//// {p1: false, p2: 0, p3: false}
	//
	//mutatePtr(&myStr)
	//
	//fmt.Println(myStr) // 3 ???
	//// {p1: true, p2: 666, p3: false}

	fmt.Println(unsafe.Sizeof(myStr))
}

func mutatePtr(in *st) {
	//in = &st{
	//	p1: true,
	//	p2: 666,
	//	p3: false,
	//}

	fmt.Println("in mutatePtr:", in) // 2
	// {p1: true, p2: 666, p3: false}
}
