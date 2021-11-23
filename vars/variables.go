package vars

import (
	"fmt"
)

var password string

func guardar() {
	fmt.Println("Valor de password: ", password)
}

func obtenerPassword() string {
	//llamo a APIjj
	return "ricardost"
}

func ejercicio2() {
	var num int

	var boolean bool = true    // Aquí es redundante especificar el tipo
	var str string = "Ricardo" // Aquí es redundante especificar el tipo

	// var otroNum int = 20
	otroNum := 20

	fmt.Println(num, boolean, str, otroNum)
	fmt.Println("Password: ", password)
}
