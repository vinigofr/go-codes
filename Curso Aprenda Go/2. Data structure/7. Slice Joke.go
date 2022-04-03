package main

import "fmt"

func main() {
	umSlice := []int{1, 2, 3, 4, 5}

	segundoSlice := append(umSlice[0:2], umSlice[4:]...)

	// 1 2 5
	fmt.Println(segundoSlice)

	// 1 2 5 4 5
	fmt.Println(umSlice)
}