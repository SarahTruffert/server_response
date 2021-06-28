package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func maian() {

	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(10)

	scanner := bufio.NewScanner(os.Stdin)

	max := 20

	for true { // boucle infinie
		fmt.Print("Entrez votre nombre : ")
		scanner.Scan()
		nbr, err := strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println("Entrez un nombre !")
			continue
		}

		if nbr > max || nbr < 0 {
			fmt.Println("Votre nombre doit etre comrpis entre 0 et ", max, "!")
			continue
		} else if nbr == randomInt {
			fmt.Println("Bien joué !")
			break //
		} else {
			fmt.Println("Dommage")
		}
	}
}
