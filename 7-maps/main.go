package main

import "fmt"

func main() {

	salarios := map[string]int{"Vitor": 5000, "João": 4000, "Maria": 3000}
	delete(salarios, "Maria")
	salarios["Lima"] = 6000

	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é %d\n", nome, salario)
	}

	// blank identifier
	for _, salario := range salarios {
		fmt.Printf("O salário é de %d\n", salario)
	}

}
