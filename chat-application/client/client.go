package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

func gestionErreur(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	IP   = "127.0.0.01" // IP local
	PORT = "3569"       // Port utilisé
)

func main() {

	var wg sync.WaitGroup

	// Connexion au serveur
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", IP, PORT))
	gestionErreur(err)

	wg.Add(2)

	// Entrée utilisateur :
	scanner := bufio.NewScanner(os.Stdin) // création du scanner capturant une entrée utilisateur
	fmt.Print("Entrez votre pseudo : ")

	scanner.Scan()                      // lancement du scanner
	entreeUtilisateur := scanner.Text() // stockage du résultat du scanner dans une variable

	// Pseudo > 20 caractères :
	if len(entreeUtilisateur) >= 20 {
		fmt.Println("Votre nombre doit être compris entre 0 et 20 carractères")
	}

	//L'utilisateur ne peut pas prendre un pseudo déjà utilisé
	if entreeUtilisateur != entreeUtilisateur {
		fmt.Println("Ce pseudo est déja prit.")
	}

	fmt.Println(entreeUtilisateur)

	go func() { // goroutine dédiée à l'entrée utilisateur
		log.Println("entrée user")
		defer wg.Done()
		for {
			reader := bufio.NewReader(os.Stdin)
			text, err := reader.ReadString('\n')
			gestionErreur(err)

			conn.Write([]byte(text))
		}
	}()

	go func() { // goroutine dédiée à la reception des messages du serveur
		defer wg.Done()
		for {
			message, err := bufio.NewReader(conn).ReadString('\n')
			gestionErreur(err)

			//fmt.Print("serveur : " + message)
			fmt.Print(entreeUtilisateur + " : " + message)
			//log.Println("Retourne valeur client écrite")
		}
	}()

	wg.Wait()

}
