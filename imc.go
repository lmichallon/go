package main

import "fmt"

func main() {
	// prénom
	const name = "Liswag"

	// catégories d'IMC
	const maigreur = 18.5
	const normal = 25.0
	const surpoids = 30.0

	// variables poids et taille
	var poids float64 = 70.5
	var taille float64 = 1.75

	// calcul de l'IMC
	imc := poids / (taille * taille)

	// affichage de l'IMC avec 2 décimales
	fmt.Printf("Bonjour %s !\n", name)
	fmt.Printf("IMC : %.2f\n", imc)

	// Affichage de la catégorie
	if imc < maigreur {
		fmt.Println("Catégorie : Maigreur")
	} else if imc < normal {
		fmt.Println("Catégorie : Normal")
	} else if imc < surpoids {
		fmt.Println("Catégorie : Surpoids")
	} else {
		fmt.Println("Catégorie : Obésité")
	}
}
