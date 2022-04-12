package main

import "fmt"

func main() {
	fmt.Println(fatorial(3), " <-- total")
}

func fatorial(x int) int {
	if x == 1 {
		return x
	} else {
		fmt.Println("Dentro do else deu ", x)  
		return x * fatorial(x - 1)
	}
}