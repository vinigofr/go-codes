package main

import "fmt"

// Todos os numeros inteiros positivos suportados pela CPU:
// Sao atribuidos de acordo com a arquitetura do CPU:
// Se for de 64 bits, o inteiro sera de 64bits
var inteiro8bits uint8
var inteiro16bits uint16
var inteiro32bits uint32
var inteiro64bits uint64

// Para inteiros negativos, use-se a nomenclatura de intNUMBER:
// int8, int16, int32, int64. Seus valores podem chegar a negativos onde 1 bit eh usado
// para atribuicao do sinal de "-"

// Em go, temos apelidos para uint e int 
// byte = apelido para uint8
// rune = apelido para int32

func main() {
  a := "e" // 1 byte
  b := "é" // 2 bytes
  c := "😃" // 4 bytes
  fmt.Printf("%v %v %v\n\n", a, b, c)

  d := []byte(a)
  e := []byte(b)
  f := []byte(c)
  fmt.Printf("%v %v %v", d, e, f)
}