package main

import (
	"fmt"
)

func main() {
	var age int
	fmt.Println("Qual a sua IDADE?")
	fmt.Scanln(&age)

	if age >= 18 {
		fmt.Println("Você é de maior de idade.")
	} else {
		fmt.Println("Você é menor de idade.")
	}

	if len("") > 0 {

	}
}
