package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	fmt.Println("el valor de x es: ", x)
	if x < 0 {
		fmt.Println("entró al if")
		return sqrt(-x) + "i"
	} else {
		fmt.Println("no entró al if")
		return fmt.Sprint(math.Sqrt(x))
	}
}

func recursion(x float64) string {
	fmt.Println("el valor de x es: ", x)
	if -4 < 0 {
		fmt.Println("entró al if")
		return sqrt(-x) + "i"
	} else {
		fmt.Println("no entró al if")
		return fmt.Sprint(math.Sqrt(x))
	}
}

func main() {
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-4))
}
