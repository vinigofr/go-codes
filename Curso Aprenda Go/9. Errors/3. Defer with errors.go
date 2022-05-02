package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Caso um erro ocorrer após a linha de um defer, o defer será efetivado.
func main() {
	file, err := os.Open("names.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	file.Close()

	byteString, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(byteString))
}