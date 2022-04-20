package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Compilação crusada, rodando nesse momento em um sistema", runtime.GOOS)

	fmt.Println("Mais informações sobre")
	fmt.Println("Arquiterura:", runtime.GOARCH)
	fmt.Println("Número de núcleos do processador:", runtime.NumCPU())
}