package main

import "fmt"

func main() {
	arrayDeInteiros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 23, 12, 48, 4984, 64, 56465}

	agoraDeuBom := findEven(sumAll, arrayDeInteiros...)
	fmt.Println(agoraDeuBom)

}

func sumAll(numbers ...int) int {
	total := 0

	for _, v := range numbers {
		total += v
	}
	return total
}

func findEven(f func(x ...int) int, ints ...int) int {
	sliceOfInts := []int{}
	for _, v := range ints {
		if v%2 != 0 {
			sliceOfInts = append(sliceOfInts, v)
		}
	}

	total := f(sliceOfInts...)
	return total
}
