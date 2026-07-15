package main

import (
	"fmt"
)

func dobrar(number int) (int, int, error) {
	if number < 0 {
		return 0, 0, fmt.Errorf("número negativo: %d", number)
	}
	return number * 2, number * 2, nil
}

func main() {
	result, _, err := dobrar(-10)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Resultado:", result)
}
