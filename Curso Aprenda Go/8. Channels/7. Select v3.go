package main

import "fmt"


func main() {
	par := make(chan int)
	impar := make(chan int)
	quit := make(chan bool)

	go sendNumbers(par, impar, quit)
	receive(par, impar, quit)
}

func sendNumbers(par chan int, impar chan int, quit chan bool) {
	total := 10
	for index := 1; index <= total; index++ {
		if index%2 == 0 {
			par <- index
		} else {
			impar <- index
		}
	}
	quit <- true
	close(par)
	close(impar)
}

func receive(par chan int, impar chan int, quit chan bool) {
	for {
		select {
		case v := <-par:
			fmt.Println("O número", v, "é par")
		case v := <-impar:
			fmt.Println("O número", v, "é impar")
		case <-quit:
			return
		}
	}
}
