package main

import "fmt"

func main() {
	x:= 387

	// ex1
	func(x int) {
		fmt.Println(x)
	}(x) // Auto callable

	// ex2
	exemplo2 := func(x int) {
		fmt.Println(x)
	}
	exemplo2(x)
}