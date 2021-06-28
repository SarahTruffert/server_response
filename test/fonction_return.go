package main

import "fmt"

func maxNbr(a int, b int) int {
	if a > b {
		return a
	}
	return b

}

func maien() {
	max := maxNbr(10, 30) // On stocke dans la variable max
	fmt.Println(max)

	fmt.Println("Valeur : %d , Type : %T\n", maxNbr(50, 30), maxNbr(50, 30))
}
