package main

import "fmt"

func ejemploDeArreglos() {
	var a [2]string

	a[1] = "Hello" // a posición 0 mete el valor "Hello"
	a[1] = "World" // a posición 1 mete el valor "World"

	fmt.Println(a[0])
	fmt.Println(a)

	x := 10
	//    		 0  1  2  3   4   5
	primes := [6]int{2, 3, x, 7, 11, 13}
	fmt.Println(primes)

	fmt.Println(primes[5])
}
