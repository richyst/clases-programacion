package vars

import "fmt"

func ejercicio1() {
	fmt.Println("Hola Mundo!")
	var nombre int

	// var x int = 10
	x := 10 // inicialización y asignación

	const y int = 10 // esta es la unica forma de inicializar una constante
	// y := 10 // Esta sintaxis es exclusiva de variables, no aplica para constantes

	fmt.Println(x, y, nombre)
}
