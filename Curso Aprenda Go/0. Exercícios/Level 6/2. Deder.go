package main

import "fmt"

func main() {
	fmt.Println("Primeiro");
	defer fmt.Println("Último");
	fmt.Println("Segundo");
}