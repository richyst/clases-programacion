package main

import "fmt"

type Animal struct {
	Nombre   string
	Ojos     int
	Patas    int
	Mamifero bool
	Hermoso  bool
}

func acumularNombres(arreglo []Animal) string { // debe salir "Perro, Gato"
	// Van a recibitr un arreglo de animales
	// iteran sobre el arreglo
	// obtienen el nombre de cada animal
	// fusionen los nombres de los animales "nombre1, nombre2"
	var acumulado string // acumulado = ""

	// 1.- i = 0, acumulado = "" llega arreglo[i].Nombre = "Perro"
	// 2.- i = 1, acumulado = "" llega arreglo[i].Nombre = "Perro"
	for i := 0; i < 2; i++ {
		// 1.- "" + "Perro"
		// 2.- "Perro" + "Gato"
		acumulado = acumulado + arreglo[i].Nombre
	}

	return acumulado
}

func acumularPatas(arreglo []Animal) int { // debe salir 8
	// Van a recibitr un arreglo de animales
	// iteran sobre el arreglo
	// obtienen el numero de patas de cada animal
	// sumen

	return 0
}

func main() {
	animalfavorito := Animal{
		Nombre:   "Perro",
		Ojos:     2,
		Patas:    4,
		Mamifero: true,
		Hermoso:  true,
	}

	animal2 := Animal{
		Nombre:   "Gato",
		Ojos:     2,
		Patas:    4,
		Mamifero: true,
		Hermoso:  false,
	}

	arreglo := []Animal{animalfavorito, animal2}

	fmt.Println("Nombres: ", acumularNombres(arreglo))
	fmt.Println("Patas: ", acumularPatas(arreglo))
}
