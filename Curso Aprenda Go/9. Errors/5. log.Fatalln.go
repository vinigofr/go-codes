// Alguns tipos de erros permitem que você continuem usando
// a aplicação sem problemas. Porém, há outros erros que são
// excepcionais e precisam interromper a execução do programa.

package main

import (
	"fmt"
	"log"
)

// Um desses exemplos é a função Fatalln() do pkg log
// Ele indica a necessidade de finalização do programa e entrega um timestamp
// E ele mesmo se encarrega de finalizar o programa.
// Funciona da mesma forma que o os.Exit(1)

// Como qualquer log, eu posso também colocar essa saída em um SetOutput com io.Writer
func main() {
	log.Fatalln("Algum erro aqui")
	fmt.Println(1111111)
}