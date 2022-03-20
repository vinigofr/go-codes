package main

import "fmt"

var aValue int = 10

func main() {
	// Demonstre o valor em decimal, binário e hex
	fmt.Printf("%d, %b, %x\n", aValue, aValue, aValue)

	fmt.Printf("Antes da conversão -> %b\n", aValue)
	anotherValue := aValue << 1
	fmt.Printf("%b", anotherValue)
}