package main

import "fmt"

func main() {

	// tuperrofavorito := "Maxito"
	// mequieres := true
	// numerofav := 10

	var arreglo [5]int

	arreglo[0] = 10
	arreglo[1] = 20
	arreglo[2] = 30
	arreglo[3] = 40
	arreglo[4] = 50

	// for i := 0; i < 4; i++ { // Recorrer arreglo de menos a mas
	// 	fmt.Println(arreglo[i])
	// }

	// for i := 3; i >= 0; i-- { // Recorrer arreglo de mas a menos
	// 	fmt.Println(arreglo[i])
	// }

	fmt.Println("El largo del arreglo es: ", len(arreglo))
	// for i := 0; i < len(arreglo); i++ {
	// 	fmt.Println(arreglo[i])
	// }

	for i := len(arreglo) - 1; i >= 0; i-- {
		fmt.Println(arreglo[i])
	}
}
