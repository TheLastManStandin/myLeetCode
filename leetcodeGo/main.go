// Go ----------------
package main

import (
	"fmt"
	"unsafe"
)

type backet struct {
	a bool
	b int
	c bool
}

func main() {
	fmt.Println(unsafe.Sizeof(backet{}))
}
