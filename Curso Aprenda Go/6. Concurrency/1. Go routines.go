package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
func main() {
	// Ok, quando utilizamos o "go" antes de qualquer função, ela será executada concorrentemente.
	// Isso significa que se você atirar, não poderá mais alterar o estado da munição, pois você não tem
	// Mais controle sobre ela.

	// Como a função main() termina a execução antes da hora, não dando espaço para as goroutines.

	// A solução para isso é utilizar o sync.WaitGroup()
	wg.Add(2)
	go func1()
	go func2()
	wg.Wait()
}

func func1() {
	for i := 0; i < 3; i++ {
		fmt.Println("Função 1", i)
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}	

func func2() {
	for i := 0; i < 3; i++ {
		fmt.Println("Função 2", i)
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}