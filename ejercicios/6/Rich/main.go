package main

import "fmt"

func main() {
	nombres := []string{"Ricardo", "George", "Nat", "Deya", "Nat", "Deya", "Ricardo", "George", "Nat", "Deya", "Nat", "Deya"}

	for posicion, nombre := range nombres {
		fmt.Printf("Posición: %d, Nombre: %s\n", posicion, nombre)
	}

	patas := map[string]int{
		"Perro":   4,
		"Pato":    2,
		"Araña":   8,
		"Alacrán": 6,
		"oruga":   23234234,
	}

	for animal, numero := range patas {
		fmt.Printf("Animal %s tiene %d patas\n", animal, numero)
	}
}
