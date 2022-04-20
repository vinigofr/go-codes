package main

import "fmt"

// Um pouco de "comma ok" em canais
func main() {
	channel := make(chan int)

	go func() {
		channel <- 50
		channel <- 0
		close(channel)
	}()

	v, ok := <-channel
	fmt.Println("Valor", v) // 50
	fmt.Println("Ok?", ok) // true

	fmt.Println()

	v, ok = <-channel
	fmt.Println("Valor", v) // 0
	fmt.Println("Ok?", ok) // true

	fmt.Println()

	v, ok = <-channel
	fmt.Println("Valor", v) // 0
	fmt.Println("Ok?", ok) // false
}