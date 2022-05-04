package main

import "fmt"

func main() {
	x := Soma(1, 2, 3, 4, 5)
	fmt.Println(x)
}

func Soma(numbers ...int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}

	return total
}