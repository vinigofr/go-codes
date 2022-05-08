package main

import (
	"fmt"
	"testing"
)

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

func BenchmarkSoma(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		Soma(1,2,3,4,5,6,7,8,9,10,10,11,4,84,894,984,654,984,65,465468,489,489,498,498,489,49,465,465,4654,654)
	}
}