package main

import "fmt"

var nombre string
var edad int
var vacunado bool

type Animal struct {
	depredador bool
	nombre     string
	habitat    string
	años       int
	tipo       string
}

func print(imp string) {
	fmt.Println(imp)
}

func suma(num1 int, num2 int) int {
	return num1 + num2
}

func main() {

	nombre = "Jorge"
	edad = 26
	vacunado = true

	nombre2 := &nombre

	*nombre2 = "Mariano"

	aguila := Animal{
		depredador: true,
		nombre:     "Águila",
		habitat:    "Estadio Azteca",
		años:       110,
		tipo:       "Las poderosísimas Águilas del América",
	}

	animalEspiritual := Animal{
		depredador: true,
		nombre:     "Oso Panda",
		habitat:    "China",
		años:       5,
		tipo:       "Adorbale pero letal",
	}

	arregloAnimales2 := [2]Animal{aguila, animalEspiritual}

	var arregloAnimales [2]Animal

	arregloAnimales[0] = aguila
	arregloAnimales[1] = animalEspiritual

	for i := 0; i < 2; i++ {
		fmt.Println(sum2)
	}

	//fmt.Println(arregloAnimales)
}
