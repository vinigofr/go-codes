package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 20)

	go func() {
		c <- "vinicius"
	}()


	fmt.Println(<-c)
	fmt.Println(<-c)
}
