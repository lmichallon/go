package main

import (
	"errors"
	"fmt"
)

// operer effectue l'opération demandée sur a et b
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

// retourne une closure pour l'opération donnée
func creerOperation(operation string) func(float64, float64) float64 {
	return func(a, b float64) float64 {
		resultat, _ := operer(a, b, operation)
		return resultat
	}
}

func main() {
	fmt.Println("entrez : nombre opération nombre (ex: 10 5 +)")
	fmt.Println("tapez 'quit' pour quitter.")

	for {
		var a, b float64
		var operation string

		fmt.Scan(&a, &b, &operation)

		if operation == "quit" {
			fmt.Println("au revoir !")
			break
		}

		resultat, err := operer(a, b, operation)
		if err != nil {
			fmt.Println("erreur :", err)
		} else {
			fmt.Printf("%.2f %s %.2f = %.2f\n", a, operation, b, resultat)
		}
	}
}
