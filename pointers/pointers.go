package main

import "fmt"

func main() {
	datoOriginal := 1
	otroDato := 3
	fmt.Println(datoOriginal)

	enlace := &datoOriginal // hago referencia a la ubicación en memoria pero NO al valor
	fmt.Println("Primera ubicación en memoria: ", enlace)
	*enlace = 20
	fmt.Println("Nuevo Valor de DatoOriginal: ", datoOriginal)
	//----------------------------------------------
	enlace = &otroDato
	fmt.Println("Segunda ubicación en memoria: ", enlace)
	*enlace = 60
	fmt.Println("Nuevo Valor de otroDato: ", otroDato)
}
