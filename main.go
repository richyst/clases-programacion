package main

import "fmt"

var password string

func guardar() {
	fmt.Println("Valor de password: ", password)
}

func obtenerPassword() string {
	//llamo a API
	return "ricardost"
}

func main() {
	guardar()
	password = obtenerPassword()
	guardar()
}
