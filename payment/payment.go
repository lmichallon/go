package main

import (
	"errors"
	"fmt"
	"math"
)

// 1. Interface Payeur
type Payeur interface {
	Payer(montant float64) (string, error)
}

// 2. CarteCredit
type CarteCredit struct {
	Numero    string
	Titulaire string
	Solde     float64
}

func (c *CarteCredit) Payer(montant float64) (string, error) {
	if montant > c.Solde {
		return "", errors.New("solde insuffisant sur la carte")
	}
	c.Solde -= montant

	// on n'affiche que les 4 derniers chiffres, comme un vrai reçu
	dernier4 := c.Numero
	if len(c.Numero) >= 4 {
		dernier4 = c.Numero[len(c.Numero)-4:]
	}
	return fmt.Sprintf("Transaction CB #%s confirmée", dernier4), nil
}

// 3. PayPal
type PayPal struct {
	Email string
	Solde float64
}

func (p *PayPal) Payer(montant float64) (string, error) {
	if montant > p.Solde {
		return "", errors.New("solde PayPal insuffisant")
	}
	p.Solde -= montant
	return fmt.Sprintf("Paiement PayPal de %.2f€ vers %s", montant, p.Email), nil
}

// 4. Crypto 
const tauxBTC = 50000.0 // 1 BTC = 50000€

type Crypto struct {
	Adresse string
	Solde   float64 // solde exprimé en euros
	Monnaie string
}

func (cr *Crypto) Payer(montant float64) (string, error) {
	if montant > cr.Solde {
		return "", errors.New("solde insuffisant pour la conversion")
	}
	cr.Solde -= montant

	quantite := math.Round(montant/tauxBTC*1000) / 1000
	return fmt.Sprintf("Paiement de %.3f %s (%.2f€) vers %s", quantite, cr.Monnaie, montant, cr.Adresse), nil
}

// Vérifications statiques à la compilation
var _ Payeur = &CarteCredit{}
var _ Payeur = &PayPal{}
var _ Payeur = &Crypto{}

// 5. ProcesserPanier
func ProcesserPanier(payeur Payeur, articles []float64) {
	total := 0.0
	for _, prix := range articles {
		total += prix
	}
	fmt.Printf("Total du panier : %.2f€\n", total)

	// type switch pour identifier le mode de paiement utilisé
	switch p := payeur.(type) {
	case *CarteCredit:
		fmt.Println("Mode de paiement : Carte de crédit (", p.Titulaire, ")")
	case *PayPal:
		fmt.Println("Mode de paiement : PayPal (", p.Email, ")")
	case *Crypto:
		fmt.Println("Mode de paiement : Crypto-monnaie (", p.Monnaie, ")")
	default:
		fmt.Println("Mode de paiement : inconnu")
	}

	confirmation, err := payeur.Payer(total)
	if err != nil {
		fmt.Println("Erreur de paiement :", err)
		return
	}
	fmt.Println(confirmation)
}

func main() {
	articles := []float64{29.99, 15.50, 120.00}

	cb := &CarteCredit{Numero: "4111111111111234", Titulaire: "Alice Martin", Solde: 500}
	ProcesserPanier(cb, articles)
	fmt.Println()

	pp := &PayPal{Email: "bob@example.com", Solde: 200}
	ProcesserPanier(pp, articles)
	fmt.Println()

	crypto := &Crypto{Adresse: "1A2b3C4d5E...", Solde: 1000, Monnaie: "BTC"}
	ProcesserPanier(crypto, articles)
	fmt.Println()

	// démonstration de l'erreur en cas de solde insuffisant
	cbPauvre := &CarteCredit{Numero: "5555666677778888", Titulaire: "Charlie", Solde: 10}
	ProcesserPanier(cbPauvre, articles)
}