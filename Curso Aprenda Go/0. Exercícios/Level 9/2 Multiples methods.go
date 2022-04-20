package main

import "fmt"

type Pessoa struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

type Humanos interface {
	Falar()
}

func (p *Pessoa) Falar() {
	fmt.Println("Oi, meu nome é", (*p).Nome);
}

func tellMeSomething(h Humanos) {
	h.Falar()
}

func main() {
	pessoa1 := Pessoa{
		Nome: "Vinicius",
		Idade: 22,
	}

	pessoa1.Falar()
	tellMeSomething(&pessoa1)
	
}