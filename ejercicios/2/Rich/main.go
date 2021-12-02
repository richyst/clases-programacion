package main

import "fmt"

var texto string
var num int
var booleano bool
var pointer *string

type Animal struct {
	Nombre   string
	Ojos     int
	Patas    int
	Mamifero bool
	Hermoso  bool
}

func main() {
	texto = "Ricardo"
	num = 10
	booleano = true

	pointer = &texto

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

	arregloAnimales := [2]Animal{animalfavorito, animal2}

	for i := 0; i < 2; i++ {
		fmt.Println(arregloAnimales[i].Nombre)
	}

}
