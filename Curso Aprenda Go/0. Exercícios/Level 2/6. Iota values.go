package main

import "fmt"

const (
	_ = iota + 2022
	nextYear
	nextYear2
	nextYear3
	nextYear4
)

func main() {
	fmt.Printf("Os próximos anos serão %d, %d, %d e %d",
		nextYear, nextYear2, nextYear3, nextYear4)
}