package main

import "fmt"

type Coche struct {
	Marca  string
	Modelo string
	Status string
}

func arreglarCocheMal(carrito Coche) { // paso por valor
	carrito.Status = "Arreglado" // aqui modifica una copia del coche original y no toca el original
	fmt.Println("Coche recibido: ", carrito)
}

func arreglarCocheBien(carrito *Coche) { // paso por referencia
	carrito.Status = "Arreglado" // aquí si modifica el coche original no una copia
	fmt.Println("Coche recibido: ", carrito)
}

func main() {
	miCoche := Coche{
		Marca:  "Ford",
		Modelo: "Fiesta",
		Status: "Descompuesto",
	}

	fmt.Println("Coche antes de arreglar: ", miCoche)

	// arreglarCocheMal(miCoche) // estoy mandando a la función una copia de mi coche pero no es el coche original

	arreglarCocheBien(&miCoche) // estoy mandando a la función la ubicación en memoria del coche original

	fmt.Println("Coche despues de arreglar: ", miCoche)
}
