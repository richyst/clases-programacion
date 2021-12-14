// go run ejercicios/4/Nat/main.go
package main

import "fmt"

// RESULTADO ESPERADO: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
func recorrerArreglo(arreglo []int) {

	for i := 0; i < 10; i++ {

		fmt.Println(arreglo[i])
	}

}

// RESULTADO ESPERADO: 1, 3, 5, 7, 9
func imprimirImpares(arreglo []int) {

	for i := 0; i < 10; i = i + 2 {

		fmt.Println(arreglo[i])
	}

}

// RESULTADO ESPERADO: 10, 9, 8, 7, 6, 5, 4, 3, 2, 1
func imprimirEnReversa(arreglo []int) {

	for i := len(arreglo) - 1; i >= 0; i-- {

		fmt.Println(arreglo[i])
	}
}

// Sacar el promedio de los valores en el arreglo
// RESULTADO ESPERADO: 6
func promedio(arreglo []int) int {

	for i 

	return 0
}

// Esta función recibe un arreglo, deben regresar un arreglo igual pero no pueden modificar las lineas que ya están
// RESULTADO ESPERADO: [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
func copiarArreglo(arreglo []int) []int {
	nuevoArreglo := []int{} // NO MODIFICAR ESTA LINEA

	for i := 0; i < len(arreglo); i++ {
		nuevoArreglo[i] = arreglo[i]
	}
	return nuevoArreglo // NO MODIFICAR ESTA LINEA
}

// RESULTADO ESPERADO: 10
func encontrarElMayor(arreglo []int) int {

	return 0
}

// RESULTADO ESPERADO: []int{}
func encontrarRepetidos(arreglo []int) []int {
	return []int{}
}

func main() {

	arreglo := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// recorrerArreglo(arreglo)
	// imprimirImpares(arreglo)
	// imprimirEnReversa(arreglo)
	// fmt.Println(promedio(arreglo))
	fmt.Println(copiarArreglo(arreglo))
	// fmt.Println(encontrarElMayor(arreglo))
	// fmt.Println(encontrarRepetidos(arreglo))
}
