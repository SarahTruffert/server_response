/* package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin) // création du scanner capturant une entrée utilisateur
    fmt.Print("Entrez quelque chose : ")
    scanner.Scan()                      // lancement du scanner
    entreeUtilisateur := scanner.Text() // stockage du résultat du scanner dans une variable
    fmt.Println(entreeUtilisateur)
}

*/

/* package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Entrez un nombre entier : ")
	scanner.Scan()

	nbr, _ := strconv.Atoi(scanner.Text()) // Pour convertir type en str

	fmt.Printf("res : %d\n", (nbr + 6))
}

*/

/*
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Entrez votre age : ")
	scanner.Scan()
	age, err := strconv.Atoi(scanner.Text())
	if err != nil {
		// Si la conversion n'a pas marché on affiche un message d' erreur.
		fmt.Println("UN ENTIER")
		// Et on quitte !
		os.Exit(2) // Quitte
	}

	if age < 17 {
		fmt.Println("Dehor !")
	} else {
		fmt.Println("Entrez !")
	}

}
*/

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func mainy() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Entrez votre age : ")
	scanner.Scan()
	age, err := strconv.Atoi(scanner.Text())
	if err != nil {
		// Si la conversion n'a pas marché on affiche un message d' erreur.
		fmt.Println("UN ENTIER")
		// Et on quitte !
		os.Exit(2) // Quitte
	}

	fmt.Print("Entrez votre prenom : ")
	scanner.Scan()
	prenom := scanner.Text()

	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(2) // Retourne aléatoirement
	randomInt2 := rand.Intn(2)

	if age < 18 {
		fmt.Println("Dehor !")
	} else if prenom == "Roro" || prenom == "roro" {
		fmt.Println("Ah c'est toi Roro dehors !")
	} else if age == 18 && randomInt == 1 { // Si le client est chanceux
		fmt.Println("Ok entre je suis de bonne humeur")
	} else if randomInt2 == 0 {
		fmt.Println("Pas de chance, dehors !")
	} else {
		fmt.Println("Entrez :)")
	}

}
