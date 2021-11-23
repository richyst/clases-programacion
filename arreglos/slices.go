package main

import "fmt"

func slices() {
	names := [4]string{
		"John",   // 0
		"Paul",   // 1
		"George", // 2
		"Ringo",  // 3
	}
	fmt.Println(names)

	a := names[0:2]
	fmt.Println("A: ", a)

	b := names[1:3] // [0: Paul, 1: George]
	fmt.Println("B: ", b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	todosMenosElUltimo := names[1:4] // [Paul, George, Ringo]
	todosMenosElUltimo[2] = "Natalia"

	fmt.Println("Todos menos el ultimo: ", todosMenosElUltimo)
	fmt.Println("Names: ", names)

	numero := 10
	fmt.Println("Ubicacion de memoria de numero: ", &numero)
}
