package main

import "fmt"

type MemRepo struct {
	Data string
}

type Repository interface{}

func CreateRepoA() Repository {
	var ptr *MemRepo
	return ptr
}

func CreateRepoB() Repository {
	return nil
}

func main() {
	fmt.Println(CreateRepoA() == CreateRepoB())
}
