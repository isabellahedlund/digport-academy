package main

import (
	"fmt"
)

func main() {
	//declare a variavel numero:
	var numero int
	// pergunte ao usuario um numero inteiro:
	fmt.Println("Digite um número inteiro: ")
	// leia o valor e atribua a variavel 'numero'
	fmt.Scanf("%d", &numero)

	if numero > 0 {
		fmt.Println("positivo!")
	} else if numero < 0 {
		fmt.Println("negativo!")
	} else {
		fmt.Println("zero!")
	}

}
