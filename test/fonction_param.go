package main

import (
	"fmt"
)

func afichage(nom string, age int) {
	fmt.Println("Bonjour", nom, "vous avez", age, "ans")
}

func metn() {
	affichage("Hatim", 9)
	affichage("Alex", 12)
	affichage("Clara", 35)
}
