package main

import "fmt"

func fn1() {
	var texto string = "Ricardo Salcedo Trejo"

	var apuntador *string = &texto

	fmt.Println(texto)      // imprime el valor
	fmt.Println(apuntador)  // imprime la dirección en memoria
	fmt.Println(*apuntador) // imprime la dereferencia de la direccion en memoria que es igual al valor

	*apuntador = "Natalia Gonzalez Hdz"

	fmt.Println(texto)
}
