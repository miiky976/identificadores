// miiky976
// Usando nombres feos como siempre B)
// -----------------------------------

package main

import (
	"fmt"
	"regexp"
)

var (
	letters string = "^[a-zA-Z]+$"
	numnlet string = "^[a-zA-Z0-9]+$"
	others  string = "^[*$#%&/]+$"
)

type tokenizer interface {
	Start()
}

type Token struct {
	input string
	chain string
}

func (token *Token) Start() {
	// Estado inicial
	for {
		fmt.Print("Estado inicial---Ingresa un caracter: ")
		fmt.Scanf("%s", &token.input)
		q0, _ := regexp.MatchString(letters, token.input)
		if q0 {
			token.chain += token.input
			break
		} else {
			fmt.Println("Token no valido")
		}
	}

	for {
		fmt.Print("Estado en q1-----Ingresa un caracter: ")
		fmt.Scanf("%s", &token.input)
		q1, _ := regexp.MatchString(numnlet, token.input)
		q2, _ := regexp.MatchString(others, token.input)
		if q1 {
			token.chain += token.input
		} else if q2 {
			fmt.Println("Estado en q2---- que tenga un buen dia")
			break
		}
	}

	fmt.Println("Cadena resultante:", token.chain)
}

func main() {
	identificador := new(Token)
	identificador.Start()
}
