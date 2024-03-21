package main

import (
	"fmt"
)

func main() {

	var b int
	a := 4

	fmt.Printf("Digite um número: ")
	fmt.Scanf("%d", &b)
	soma := a + b

	fmt.Println("My Return:", soma)
}
