package main

import "fmt"

func main() {
	arreglo := make(map[int]string) // NO es un arreglo, esto es un mapa que se comporta como arreglo
	arreglo[1] = "Ricardo"
	arreglo[2] = "Nat"
	arreglo[3] = "Llorch"
	arreglo[4] = "Deya"

	// persona1 := make(map[string]string)
	// persona1["Nombre"] = "Ricardo"
	// persona1["Apellido"] = "Salcedo"
	// persona1["Edad"] = "27"
	persona1 := map[string]string{
		"Nombre":   "Ricardo",
		"Apellido": "Salcedo",
		"Edad":     "27",
	}

	persona2 := make(map[string]string)
	persona2["Nombre"] = "Nat"
	persona2["Apellido"] = "Gonzalez"
	persona2["Edad"] = "26"

	persona3 := make(map[string]string)
	persona3["Nombre"] = "Llorch"
	persona3["Apellido"] = "Martos"
	persona3["Edad"] = "26"

	persona4 := make(map[string]string)
	persona4["Nombre"] = "Deya"
	persona4["Apellido"] = "Rivera"
	persona4["Edad"] = "27"

	viajantes := []map[string]string{persona1, persona2, persona3, persona4}
	fmt.Println("Grupo de Viajantes: ", viajantes) // imprime el arreglo 'viajantes'
	fmt.Println()
	fmt.Println()
	fmt.Println()

	for i := 0; i < 4; i++ {
		fmt.Printf(
			"Viajante completo: %v\nNombre completo de Viajante: %s %s\nEdad de Viajante: %s\n\n\n",
			viajantes[i],           // me regresa un elemento del arreglo de maps que es un map completo
			viajantes[i]["Nombre"], // me regresa la propiedad Nombre del map en el arreglo de viajantes
			viajantes[i]["Apellido"],
			viajantes[i]["Edad"], // me regresa la propiedad Edad del map en el arreglo de viajantes
		)
	}
	// fmt.Println("Viajantes: ", viajantes)

}
