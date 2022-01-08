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

	// recorreremos el "arreglo" ✅
	// - agarrar el elemento en la posición "i" del arreglo (nunca se debe hacer con la ultima posición)
	// - - ese elemento lo debes comparar contra todos los elementos que le siguen
	// - - si encuentras uno repetido -> lo guardas en el arreglo "numerosRepetidos"
	// regresas el arreglo "numerosRepetidos"

	numerosRepetidos := []int{}

	for i := 0; i < len(arreglo); i++ { // voy a recorrer arreglo

		// repetidos := arreglo[i]
		// fmt.Println(i)
		for j := i; j < len(arreglo); j++ {
			// if numerosRepetidos ==
			numeroRepetidos = append(numerosRepetidos, repetidos)
			fmt.Println("num; ", arreglo[i], " vs num: ", arreglo[j])
		}
		// if numerosRepetidos == arreglo[i]{
		// 	numerosRepetidos = arreglo[i]
		// }

	}

	return numerosRepetidos
}

// {1 ,  2,  3,  4,  5}
// {5 ,  6,  7,  8,  9}
// {11, 12, 13, 14, 15}
// {16, 17, 18, 19, 20}
// {21, 22, 23, 24, 25}
func imprimirArregloBi(arreglo [][]int) {
	for i := 0; i < len(arreglo); i++ {
		for j := 0; j < len(arreglo[i]); j++ {
			fmt.Println(arreglo[i][j]) // 1: 1
		}
		fmt.Println("Fin del for de j")
	}
	fmt.Println("Fin del for de i")
}

func imprimirEstudiantes(escuela [][]string) {
	// nombres := ""
	for i := 0; i < len(escuela); i++ { // itera cada salon de la escuela
		fmt.Printf("\n\nSalón %d = %v\n", i, escuela[i])
		for j := 0; j < len(escuela[i]); j++ { // itera cada estudiante de cada salón de la escuela
			fmt.Printf("Salon: %d, Alumno: %d = %s\n", i, j, escuela[i][j])
			// nombres = fmt.Sprintf("%s\n%s", nombres, escuela[i][j])
		}
	}
	// return nombres
}

func main() {
	arreglo := []int{19, 22, 31, 45, 57, 16, 76, 81, 9, 10}
	// arreglo_bi := [][]int{{1, 2, 3, 4, 5}, {5, 6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}
	// madrid := [][]string{
	// 	{"George", "nat", "Rich"}, // salon 0
	// 	{"Deya", "Jime", "Dam"},   // salon 1
	// 	{"Luis", "Jan", "Raquel"}, // salon 2
	// }
	// olinca := [][]string{ // escuelta
	// 	{ // salon 0 de olinca
	// 		"Deya", // estudiante 0 del salon 0 de olinca
	// 		"Jime", // estudiante 1 del salon 0 de olinca
	// 		"Dam",  // estudiante 2 del salon 0 de olinca
	// 	},
	// 	{"George", "nat", "Rich"}, // salon 1
	// 	{"Luis", "Jan", "Raquel"}, // salon 2
	// }
	// sagrado := [][]string{}

	// imprimirEstudiantes(madrid)
	// imprimirEstudiantes(sagrado)
	// imprimirEstudiantes(olinca)

	// var arreglo []int = []int{1,2,3,4,5,6,7,8,9,10}
	// recorrerArreglo(arreglo)
	// imprimirImpares(arreglo)
	// imprimirEnReversa(arreglo)
	// fmt.Println(promedio(arreglo))
	// fmt.Println(copiarArreglo(arreglo))
	// fmt.Println(encontrarElMayor(arreglo))
	fmt.Println(encontrarRepetidos(arreglo))
}
