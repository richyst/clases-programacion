package main

import "fmt"

var nombre string
var numero int
var booleano bool

var apuntador *string

type Animal struct {
	Nombre   string
	Ojos     int
	Patas    int
	Mamifero bool
	Hermoso  bool
}

func main() {

	nombre = "Maxito"
	numero = 10
	booleano = true

	apuntador = &nombre

	*apuntador = "Maximiliano"

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

	var arreglos [2]Animal

	arreglos[0] = animalfavorito
	arreglos[1] = animal2

	arregloscorto := [2]Animal{animalfavorito, animal2}

}

func imprimir(bonjour string) {

	fmt.Println(bonjour)
}

func suma(numero1 int, numero2 int) int {
	return numero1 + numero2
}
