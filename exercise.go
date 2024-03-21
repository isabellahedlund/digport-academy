package main

import (
	"fmt"
)

func main() {
	var name string

	fmt.Println("Qual seu nome?")
	fmt.Scan(&name)
	fmt.Println("Bem vinda " + name)

}
