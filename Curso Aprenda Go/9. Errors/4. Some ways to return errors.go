package main

import (
	"log"
	"os"
)

// Para melhor trabalhar com erros, podemos utilizar o pkg "log" builtin

func main() {
	file, err := os.Create("log.txt")

	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	
	// Através dessa linha, especificamos que agora
	// todas as saídas serão para a "var file"
	log.SetOutput(file)

// ==========

	anotherFile, err := os.Open("não existe.txt")
	if err != nil {
		// Aqui inserimos a primeira linha do "log"
		log.Println("Aconteceu um erro ->", err)

		// Mais uma linha aqui
		log.Println("Gouveia passou por aqui")
		
	}
	defer anotherFile.Close()
}