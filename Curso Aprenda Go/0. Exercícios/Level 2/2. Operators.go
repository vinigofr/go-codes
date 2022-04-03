package main

import "fmt"

const (
	a = 10
	b = 20
)

const (
	value1 = a == a
	value2 = a != b
	value3 = a <= a
	value4 = a < b
	value5 = b >= a
	value6 = b > a
)

func main() {
	fmt.Printf("%t\n", value1)
	fmt.Printf("%t\n", value2)
	fmt.Printf("%t\n", value3)
	fmt.Printf("%t\n", value4)
	fmt.Printf("%t\n", value5)
	fmt.Printf("%t\n", value6)
}