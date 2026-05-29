package main

import "fmt"

func main() {
	// slice de tâches
	taches := []string{
		"apprendre les slices en Go",
		"maîtriser la boucle for",
		"comprendre le fallthrough",
		"faire les exercices",
	}

	// ajout d'une tâche avec append
	taches = append(taches, "réviser avant le cours")

	fmt.Println("liste des tâches :")

	// boucle for unique : parcourt le slice et affiche le statut de chaque tâche
	for i, tache := range taches {
		priorite := i + 1

		// fallthrough : la priorité 1 cumule les labels CRITIQUE + URGENT + message
		switch {
		case priorite == 1:
			fmt.Print("CRITIQUE ")
			fallthrough
		case priorite == 2:
			fmt.Print("URGENT ")
			fallthrough
		case priorite <= 3:
			fmt.Printf("à faire aujourd'hui : %s\n", tache)
		default:
			fmt.Printf("NORMAL %s\n", tache)
		}
	}
}
