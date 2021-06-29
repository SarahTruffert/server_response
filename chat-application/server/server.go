package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func gestionErreur(err error) {
	log.Println("gestion d' erreur")
	if err != nil {
		panic(err)
	}
}

const (
	IP   = "127.0.0.01"
	PORT = "3569"
)

func read(conn net.Conn) {
	log.Println("Read custom message")
	message, err := bufio.NewReader(conn).ReadString('\n')
	gestionErreur(err)

	fmt.Print("Client:", string(message))

}

func main() {
	log.Println("Côté serveur")

	fmt.Println("Lancement du serveur ...")

	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%s", IP, PORT))
	gestionErreur(err)

	var clients []net.Conn // tableau de clients

	for {
		conn, err := ln.Accept()
		log.Println("Accepte connexion externe")
		if err == nil {
			clients = append(clients, conn) //quand un client se connecte on le rajoute à notre tableau
			log.Println("Client ajouté")
		}
		gestionErreur(err)
		fmt.Println("Un client est connecté depuis", conn.RemoteAddr())

		go func() { // création de notre goroutine quand un client est connecté
			log.Println("Création de la Goroutine")
			buf := bufio.NewReader(conn)

			for {
				name, err := buf.ReadString('\n')

				if err != nil {
					fmt.Printf("Client disconnected.\n")
					log.Println("Fin de la co du client")
					break
				}
				for _, c := range clients {
					c.Write([]byte(name)) // on envoie un message à chaque client
				}
			}
		}()
	}
}
