package main

import (
	"fmt"
)

func alterarOriginal(x *int) {
	*x = *x * 2
}

func main() {
	number := 10
	fmt.Println("Fora da função:", number)
	alterarOriginal(&number)
	fmt.Println("Fora da função:", number)
}
