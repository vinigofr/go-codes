package main

import "fmt"

func main() {
	theBool := true

	addressOfBool := &theBool

	fmt.Println(addressOfBool)
}