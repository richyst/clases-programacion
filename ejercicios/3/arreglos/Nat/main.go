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

	var acumuladodenombres string

	for i := 0; i < 2; i++ {
		acumuladodenombres = acumuladodenombres + arreglo[i].Nombre
	}

	return acumuladodenombres
}

func acumularPatas(arreglo []Animal) int { // debe salir 8
	// Van a recibitr un arreglo de animales
	// iteran sobre el arreglo
	// obtienen el numero de patas de cada animal
	// sumen

	var sumadepatas int

	for i := 0; i < 2; i++ {
		sumadepatas = sumadepatas + arreglo[i].Patas
	}

	return sumadepatas
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
		Patas:    3,
		Mamifero: true,
		Hermoso:  false,
	}

	arreglo := []Animal{animalfavorito, animal2}

	text1 := "ricardo"
	text2 := "salcedo"

	acumulado := text1 + ", " + text2
	fmt.Println(acumulado)
	fmt.Println("Nombres: ", acumularNombres(arreglo))
	fmt.Println("Patas: ", acumularPatas(arreglo))
}
