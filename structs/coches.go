package main

import "fmt"

type Coche struct {
	Marca       *Marca
	Modelo      string
	Cilindros   uint // uint no acepta numeros negativos (uint = unsigned integer)
	Año         int
	Asientos    int
	Combustible Combustible
}

type Combustible struct {
	string
}

type Marca struct {
	PaisOrigen   string
	AñoFundacion int
}

func main() {
	ford := Marca{
		AñoFundacion: 1910,
		PaisOrigen:   "EEUU",
	}

	hibrido := Combustible{"Hibrído"}

	fusion := Coche{
		Marca:       &ford,
		Modelo:      "Fusion",
		Cilindros:   6,
		Año:         2020,
		Asientos:    5,
		Combustible: hibrido,
	}

	fmt.Println("Marca de fusion: ", fusion.Marca)

	ford.PaisOrigen = "Alemania"

	fmt.Println("Marca de fusion: ", fusion.Marca)
}
