package main

import "fmt"

// 10       20        30
func suma(num1 int, num2 int) int {
	return num1 + num2
}

// 10       20        30
func resta(num1 int, num2 int) int {
	return num1 - num2
}

func division(num1 int, num2 int) float32 { // float32 es el tipo para los datos con decimales y 32 bits de precisión
	var resultado float32 = float32(num1) / float32(num2)
	return resultado
}

func divisionPrecisa(num1 int, num2 int) float64 { // float64 es el tipo para los datos con decimales y 64 bits de precisión
	var resultado float64 = float64(num1) / float64(num2)
	return resultado
}

func multiplicacion(num1 float32, num2 float32) float32 {
	return num1 * num2
}

// la función se llama convertidor,
// recibe un entero al que le pone el alias num1
// y regresa un entero y un float64
func convertidor(num1 int) (int, float64) {
	return num1, float64(num1)
}

func convertidorEspecial(num1 int) (especial1 int, especial2 float64) { // named returns
	especial1 = num1          // no tengo que usar := porque ya la inicializó arriba
	especial2 = float64(num1) // no tengo que usar := porque ya la inicializó arriba

	return // Aquí ya sabe que al regresar al flujo principal especial1 va primero y especial2 despues
}

func imprima(dato int) {
	fmt.Println(dato)
}

func main() {
	x := 10
	y := 3
	z := 30

	imprima(x)

	fmt.Println("Suma: ", suma(x, y))
	fmt.Println("Resta: ", resta(x, z))
	fmt.Println("Division: ", division(x, y))
	fmt.Println("Division más Precisa: ", divisionPrecisa(x, y))
	fmt.Println("Multiplicación: ", multiplicacion(float32(x), float32(x)))

	a := 1
	b := 2
	c := 3

	fmt.Println("Suma: ", suma(a, b))

	// 1.- Suma x + a
	// 2.- Suma (Resultado de 1) + c
	fmt.Println("Suma y resta: ", suma(suma(x, a), c))
	fmt.Println("Combinado: ", suma(resta(int(division(z, x)), a), z))

	entero, decimal := convertidor(x)
	fmt.Printf("Entero: %d, Decimal: %f", entero, decimal)

}
