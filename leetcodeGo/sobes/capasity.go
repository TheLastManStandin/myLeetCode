package main

import "fmt"

// https://sobes.tech/bank/go/livecode/550432e3-fe76-4ead-b5f9-247559a257c8

func main() {
	fruit := make([]string, 5)
	fruit[0] = "banana"
	fruit[1] = "cucumber"
	fruit[2] = "grape"
	fruit[3] = "watermelon"
	fruit[4] = "pineapple"

	xs := fruit[1:3:3]       // cucumber grape len 2 cap 2
	xs[0] = "apple"          // apple grape len 2 cap 2
	xs = append(xs, "mango") // apple grape

	fmt.Println(fruit)
}
