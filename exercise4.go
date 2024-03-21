package main

import (
	"fmt"
)

func main() {

	for contador := 10; contador > 0; contador-- {
		fmt.Println("Count down", contador)
	}

	fmt.Println("Happy new year!")

	for contador := 0; contador < 10; contador++ {
		fmt.Println("Count up", contador)
	}
	fmt.Println("Arrasou!")
}
