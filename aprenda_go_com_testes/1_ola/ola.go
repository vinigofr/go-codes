// Sim, como soube?
package main

import "fmt"

// Oi, teste
func Ola(nome, idioma string) string{
    if nome == "" {
        nome = "mundo"
    }

    return prefixoSaudacao(idioma) + nome    
}

func prefixoSaudacao(idioma string) (prefixo string){
    const (
        prefixoPT = "Olá, "
        prefixoES = "Hola, "
        prefixoFR = "Bonjour, "
    )

    switch idioma {
    case "Francês":
        prefixo = prefixoFR
    case "Espanhol":
        prefixo = prefixoES
    default:
        prefixo = prefixoPT
    }

    /*
        Quando se retorna dessa forma, sem nenhum valor
        o código usa como referência a variável 'prexixo'
        auto-declarada na declaração da função. Ela é
        existente e é tradada como retorno
    */
    return
}

func main() {
    fmt.Println(Ola("", "espanhol"))
}