package main

import "fmt"

func main() {
	// numeroPreferido := 19
	// nombrePila := "Llorch"
	// verdaderoFalso := true

	arr := [5]int{10, 20, 30, 40, 50} // arreglo
	// slice := []int{10, 20, 30} // slice

	// for i := 0; i < 4; i++ {
	// 	fmt.Println(arr[i])
	// }

	// for i := 3; i >= 0; i-- {
	// 	fmt.Println(arr[i])
	// }

	fmt.Println("El largo del arreglo es: ", len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Println(arr[i])
	}

}
