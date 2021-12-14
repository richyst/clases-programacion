package main

import "fmt"

// Posiciones          0, 1, 2, 3, 4, 5, 6, 7, 8,  9
// RESULTADO ESPERADO: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
func recorrerArreglo(arreglo []int) {
	// i := 0 <- inicialización
	// i < 10 <- condición
	// i++    <- actualización
	for i := 0; i < 10; i = i + 2 {
		fmt.Println(arreglo[i])
	}
}

// RESULTADO ESPERADO: 1, 3, 5, 7, 9
func imprimirImpares(arreglo []int) {
	for i := 0; i < 10; i = i + 2 {
		fmt.Println(arreglo[i])
	}
}

//                     0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
// RESULTADO ESPERADO: 10, 9, 8, 7, 6, 5, 4, 3, 2, 1
func imprimirEnReversa(arreglo []int) {
	fmt.Println("Largo de arreglo: ", len(arreglo))
	for i := 9; i >= 0; i-- {
		fmt.Println(arreglo[i])
	}
}

// Sacar el promedio de los valores en el arreglo
// RESULTADO ESPERADO: 6
func promedio(arreglo []int) int {
	return 0
}

// Esta función recibe un arreglo, deben regresar un arreglo igual pero no pueden modificar las lineas que ya están
// RESULTADO ESPERADO: [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
func copiarArreglo(arreglo []int) [10]int {
	nuevoArreglo := [10]int{} // NO MODIFICAR ESTA LINEA
	nuevoArreglo[0] = arreglo[0]

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
	// var arreglo []int = []int{1,2,3,4,5,6,7,8,9,10}
	// recorrerArreglo(arreglo)
	// imprimirImpares(arreglo)
	// imprimirEnReversa(arreglo)
	// fmt.Println(promedio(arreglo))
	fmt.Println(copiarArreglo(arreglo))
	// fmt.Println(encontrarElMayor(arreglo))
	// fmt.Println(encontrarRepetidos(arreglo))
}
