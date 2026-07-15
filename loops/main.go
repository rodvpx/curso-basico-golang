package main

import (
	"fmt"
)

func main1() {

	for i := 1; i <= 5; i++ {
		fmt.Printf("Essa é a interação numero %d\n", i)
	}
}

func main() {
	var age int
	for {
		fmt.Printf("Digite sua idade: (maior de 18)\n")
		_, err := fmt.Scanln(&age)
		if err != nil {
			fmt.Println("Entrada invalida (somente números interios), por favor repita")
			continue
		}

		if age < 18 {
			fmt.Println("Entrada invalida (maior de 18), por favor repita")
			continue
		}
		break
	}
	fmt.Println("Entrada autorizada!")
}
