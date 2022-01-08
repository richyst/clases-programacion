package main

import "fmt"

func main() {
	nombres := []string{"Ricardo", "George", "Nat", "Deya", "Nat", "Deya"}

	repeticiones := map[string]int{
		"Ricardo": 0,
		"George":  0,
		"Nat":     0,
		"Deya":    0,
	}

	// for i := 0; i < len(nombres); i++ {
	// 	for j := i + 1; j < len(nombres); j++ { // j = j +1
	// 		if nombres[i] == nombres[j] {
	// 			fmt.Println(nombres[i], i, j)
	// 			repeticiones[nombres[i]] = repeticiones[nombres[i]] + 1 // nombres[i] es igual a "Ricardo" "Deya" "Nat" o "George"
	// 			break
	// 		}
	// 	}
	// }

	for i, pivote := range nombres {
		for _, nombre := range nombres[i+1:] { // de nombres haz un slice a partir de i + 1 hasta el final
			if pivote == nombre {
				repeticiones[pivote]++
				break
			}
		}
	}

	fmt.Println(repeticiones)

}
