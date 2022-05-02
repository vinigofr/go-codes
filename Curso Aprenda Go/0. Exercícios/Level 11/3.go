package main

import (
	"errors"
	"fmt"
)

type erroEspecial struct {
	mensagemDeErro string
}

func (err erroEspecial) Error() error {
	return errors.New(err.mensagemDeErro)
}


func recebeErro(e error) {
	fmt.Println("Argumento recebido", e);
};

func main() {
	novaMensagem := erroEspecial{
		mensagemDeErro: "😀Deu algo um pouco ruim por aqui!",
	}

	erroRetornado := novaMensagem.Error()
	recebeErro(erroRetornado)
}

