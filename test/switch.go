package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func mainee() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Votre choix : ")
	scanner.Scan()
	choix, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Entrez un entier")
		os.Exit(2)
	}

	switch choix { // La variable à vérifier
	case 0, 1:
		println("George Boole ! ")
	case 7:
		println("William Van de Walle !")
	case 23:
		println("Jimmy still OK")
	case 42:
		println("La réponse à tout")
	case 666:
		println("Hell yeah")
	default:
		println("Bad choice")
	}
}
