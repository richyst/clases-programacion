package main

import (
	"fmt"
)

func main() {
	mono := []int{1, 2, 3} // [ 1 2 3]

	fmt.Println("Arreglo mono:", mono)

	//      | 1 2 3 |
	//      | 4 5 6 |
	//      | 7 8 9 |

	// bi := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}} // sintaxis corta y recomendada
	bi := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	fmt.Println("Arrego bi: ", bi)

	/* asdf
	Comentario acotado
	sdqd */

	tri := [][][]int{
		[][]int{
			[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}, // 0
		},
		[][]int{
			[]int{10, 20, 30} /* 0 */, []int{40 /*0*/, 50 /*1*/, 60 /*2*/} /* 1 */, []int{70, 80, 90}, /*2*/ // 1
		},
		[][]int{
			[]int{100, 200, 300}, []int{400, 500, 600}, []int{700, 800, 900}, // 2
		},
	}
	fmt.Println("Arreglo tri: ", tri)
	fmt.Println("Posición 1,1,1: ", tri[1][1][1])
}
