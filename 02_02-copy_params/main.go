package main

import (
	"fmt"
)

func alterarCopia(x int) int {
	fmt.Println("Recebido como:", x)
	x = x * 2
	fmt.Println("atualizado para:", x)
	return x
}

func main() {
	number := 10
	fmt.Println("Fora da função:", number)
	result := alterarCopia(number)
	fmt.Println("Fora da função:", result)
}
