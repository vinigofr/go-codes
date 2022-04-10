package main

import "fmt"

func main() {
	x := struct {
		nome  string
		idade int
	}{
		nome:  "Vinicius",
		idade: 50,
	}

	fmt.Println(x)
}