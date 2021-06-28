package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Args to choose 2tween .txt or input user :

func main() {
	if len(os.Args) > 1 {
		configPath := os.Args[1]
		txt(configPath)
		fmt.Println("0") // For debugging purposes
	} else {
		input()
		fmt.Println("1")
		//txt(configPath)
	}

}

// Read by the file :
func txt(configPath string) []string {
	f, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	cvide := []string{}
	for scanner.Scan() {
		jetestocke := scanner.Text()
		cvide = append(cvide, jetestocke)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// A la fin de txt close :
	//defer f.Close()
	f.Close()
	fmt.Println(cvide)
	white_rab(cvide)
	return cvide
}

// Read by input :
func input() []string {
	scanner := bufio.NewScanner(os.Stdin) // création du scanner capturant une entrée utilisateur
	fmt.Print("Entrez une position : ")
	scanner.Scan()                      // lancement du scanner
	entreeUtilisateur := scanner.Text() // stockage du résultat du scanner dans une variable
	words := strings.Fields(entreeUtilisateur)

	//fmt.Println(words, len(words))
	//fmt.Println(words)

	white_rab(words)
	//main()
	return words
}

// Déchiffrer les positions :
func white_rab(words []string) []int {
	//fmt.Println("words", words)
	listevide := []int{}
	for _, word := range words {
		//fmt.Println("word", word)
		if word == "front" {
			listevide = append(listevide, 0, 1)
		} else if word == "back" {
			listevide = append(listevide, 0, -1)
		} else if word == "right" {
			listevide = append(listevide, 1, 0)
		} else {
			fmt.Println("Other: Left")
			listevide = append(listevide, -1, 0)
			fmt.Println("liste vide", listevide)

		}
	}
	add_position_premier(listevide, 0)
	return listevide
}

// Additionner données [x + x] [y + y]:
// Compter par 2
func add_position_premier(listevide []int, index int) int {
	res := 0
	for i := index; i < len(listevide)-1; i += 2 {
		res += (listevide[i])
	}

	//resultat_final(res1)
	fmt.Println("Hey, White Rabbit, do you want to know where 2 go ?")
	fmt.Println(" OK ! GO ! ", res)
	return res
}
