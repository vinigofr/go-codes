package main

import (
	"fmt"
)


var canal1 = make(chan int)
func main() {

	go putNumbersOnChannel(canal1)

	for v := range canal1 {
		fmt.Println(v)
	}
}

func putNumbersOnChannel(canal1 chan int) {
	for i := 0; i < 10; i++ {
		canal1 <- 10
	}
	close(canal1)
}
