package main

import "fmt"

func main() {
	sum := 0

	for i := 0; i < 3; i++ {
		sum = i + 1 // 1 + 1
		fmt.Printf(
			"La suma lleva: %d, i es: %d y un texto:%s\n", // primer argumento
			sum,           // segundo argumento
			i,             // tercero argumento
			"Holi\n\tgoe", // cuarto argumento
		)
	}

	fmt.Println("Fin de iteración!\nHoli despues del enter", sum)
}
