// Sim, como soube?
package main

import "fmt"

// Oi, teste
func Ola(value string) string{
    var prefixoOlaPT = "Olá"
    if value != "" {
        return fmt.Sprintf("%v %v", prefixoOlaPT, value)
    }
    return "Olá"
}

func main() {
    fmt.Println(Ola(""))
}