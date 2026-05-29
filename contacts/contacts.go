package main

import "fmt"

// contact avec ses informations de base
type Personne struct {
	Prenom string
	Nom    string
	Age    int
	Email  string
}

// retourne le prénom et le nom concaténés
func (personne Personne) NomComplet() string {
	return fmt.Sprintf("%s %s", personne.Prenom, personne.Nom)
}

// retourne une ligne de présentation complète
func (personne Personne) Presentation() string {
	return fmt.Sprintf("%s, %d ans — %s", personne.NomComplet(), personne.Age, personne.Email)
}

// adresse postale
type Adresse struct {
	Rue        string
	Ville      string
	CodePostal string
}

// retourne l'adresse formatée sur une ligne
func (adresse Adresse) Format() string {
	return fmt.Sprintf("%s, %s %s", adresse.Rue, adresse.CodePostal, adresse.Ville)
}

// embed Personne et Adresse + ajoute infos pro
type Employe struct {
	Personne
	Adresse
	Poste   string
	Salaire float64
}

// retourne toutes les infos de l'employé
func (employe Employe) FicheEmploye() string {
	return fmt.Sprintf(
		"[Employé] %s\n  Poste   : %s\n  Salaire : %.2f €\n  Adresse : %s\n  Contact : %s",
		employe.NomComplet(), employe.Poste, employe.Salaire, employe.Adresse.Format(), employe.Email,
	)
}

// augmente le salaire d'un pourcentage donné
// pointeur pour modifier la valeur en place
func (employe *Employe) AugmenterSalaire(pct float64) {
	employe.Salaire += employe.Salaire * pct / 100
}

// embed Personne + ajoute les infos scolaires
type Etudiant struct {
	Personne
	Promo   string
	Moyenne float64
}

// retourne la mention selon la moyenne
func (etudiant Etudiant) MentionObtenue() string {
	switch {
	case etudiant.Moyenne >= 16:
		return "Très Bien"
	case etudiant.Moyenne >= 14:
		return "Bien"
	case etudiant.Moyenne >= 12:
		return "Assez Bien"
	default:
		return "Passable"
	}
}

// retourne toutes les infos de l'étudiant sous forme de texte
func (etudiant Etudiant) FicheEtudiant() string {
	return fmt.Sprintf(
		"[Étudiant] %s\n  Promo   : %s\n  Moyenne : %.2f — %s\n  Contact : %s",
		etudiant.NomComplet(), etudiant.Promo, etudiant.Moyenne, etudiant.MentionObtenue(), etudiant.Email,
	)
}

func main() {
	// slice d'employés
	employes := []Employe{
		{
			Personne: Personne{"Alice", "Martin", 32, "alice.martin@corp.fr"},
			Adresse:  Adresse{"12 rue de la Paix", "Paris", "75001"},
			Poste:    "Développeuse Go",
			Salaire:  3800,
		},
		{
			Personne: Personne{"Bob", "Dupont", 45, "bob.dupont@corp.fr"},
			Adresse:  Adresse{"8 avenue Foch", "Lyon", "69006"},
			Poste:    "Chef de projet",
			Salaire:  4500,
		},
	}

	// slice d'étudiants
	etudiants := []Etudiant{
		{Personne: Personne{"Clara", "Petit", 21, "clara.petit@esgi.fr"}, Promo: "B3 Infra", Moyenne: 15.4},
		{Personne: Personne{"Dylan", "Moreau", 20, "dylan.moreau@esgi.fr"}, Promo: "B3 Dev", Moyenne: 11.2},
	}

	fmt.Println("=== Fiches employés ===")
	// for par index pour pouvoir appeler AugmenterSalaire
	for index := range employes {
		employes[index].AugmenterSalaire(10)
		fmt.Println(employes[index].FicheEmploye())
		fmt.Println()
	}

	fmt.Println("=== Fiches étudiants ===")
	for _, etudiant := range etudiants {
		fmt.Println(etudiant.FicheEtudiant())
		fmt.Println()
	}
}
