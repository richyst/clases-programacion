package main

import "fmt"

// Persona is a struct representing people
type Persona struct {
	Nombre   string
	Apellido string
	Edad     int
	ColorFav string
	Hijos    []Persona
	Alergias []Alergia
}

type Alergia struct {
	Cura    string
	Sintoma Sintoma
}

type Sintoma struct {
	Descripcion    string
	TiempoReaccion int
}

func ejercicio() {
	persona1 := Persona{
		Nombre:   "Ricardo",
		Apellido: "Salcedo",
		Edad:     27,
		Hijos: []Persona{
			{
				Nombre:   "Ricardito",
				Apellido: "Gonzalez-Salcedo",
				Edad:     0,
			},
			{
				Nombre:   "Natalita",
				Apellido: "Gonzalez-Ex",
				Edad:     27,
				Hijos: []Persona{
					{
						Nombre:   "Natalita Jr",
						Apellido: "Gonzalez-Ex",
						Edad:     2,
					},
				},
			},
		},
	}
	persona2 := Persona{
		Nombre:   "Nat",
		Apellido: "González",
		Edad:     26,
	}
	persona3 := Persona{
		Nombre:   "Llorch",
		Apellido: "Martos",
		Edad:     26,
		Alergias: []Alergia{
			{
				Sintoma: Sintoma{
					Descripcion:    "Ganas de suicidio",
					TiempoReaccion: 1,
				},
				Cura: "Dosis intrapiernosa de mitrocin",
			},
		},
	}
	persona4 := Persona{
		Nombre:   "Deya",
		Apellido: "Rivera",
		Edad:     27,
	}

	viajantes := []Persona{persona1, persona2, persona3, persona4}

	for i := 0; i < 4; i++ {
		fmt.Printf("Viajante completo: %v\n\n\n", viajantes[i].Hijos)
	}

}
