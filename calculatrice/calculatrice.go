package main

import (
	"errors"
	"fmt"
)

// fait l'opération demandée sur a et b
// retourne une erreur si l'opération est inconnue OU en cas de division par zéro
func operer(a, b float64, operation string) (float64, error) {
	switch operation {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("division par zéro impossible")
		}
		return a / b, nil
	default:
		return 0, errors.New("opération inconnue : " + operation)
	}
}

// retourne une closure qui applique l'opération donnée à deux nombres
func creerOperation(operation string) func(float64, float64) float64 {
	return func(a, b float64) float64 {
		resultat, _ := operer(a, b, operation)
		return resultat
	}
}

func main() {
	fmt.Println("entrez : nombre opération nombre (ex: 10 + 5)")
	fmt.Println("tapez 'quit' pour quitter.")

	for {
		// lit le premier token pour détecter "quit" avant de lire les nombres
		var input string
		fmt.Scan(&input)

		// si input est "quit" on quitte la boucle
		if input == "quit" {
			fmt.Println("au revoir !")
			break
		}

		var a, b float64
		var operation string
		// utilise Sscan pour lire le nombre et laisser le reste de la ligne pour l'opération et le second nombre
		_, err := fmt.Sscan(input, &a)
		// si le premier token n'est pas un nombre, on affiche une erreur et on continue
		if err != nil {
			fmt.Println("format invalide. exemple : 10 + 5")
			continue
		}
		// lit l'opération et le second nombre
		fmt.Scan(&operation, &b)

		// effectue l'opération et affiche le résultat ou l'erreur
		resultat, err := operer(a, b, operation)
		if err != nil {
			fmt.Println("erreur :", err)
		} else {
			fmt.Printf("%.2f %s %.2f = %.2f\n", a, operation, b, resultat)
		}
	}
}
