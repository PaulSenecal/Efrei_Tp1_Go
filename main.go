/* Author: Paul Senecal
*	Date: 2023-10-01
*	Description: Programme d'annuaire
*	Usage:
*		go run main.go --action ajouter --nom "Charlie" --prenom "John" --tel "0811223344"
*		go run main.go --action rechercher --nom "Alice"
*		go run main.go --action lister
*		go run main.go --action supprimer --nom "Charlie"
*		go run main.go --action modifier --nom "Charlie" --prenom "Charles" --tel "0123456789"
 */

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/PaulSenecal/tp1/annuaire"
)

func main() {
	var action = flag.String("action", "", "Action à effectuer: ajouter, rechercher, lister, supprimer, modifier")
	var nom = flag.String("nom", "", "Nom du contact")
	var prenom = flag.String("prenom", "", "Prénom du contact")
	var tel = flag.String("tel", "", "Numéro de téléphone du contact")

	flag.Parse()

	fmt.Println("Programme Annuaire")
	fmt.Println("==================")

	switch *action {
	case "ajouter":
		if *nom == "" || *prenom == "" || *tel == "" {
			fmt.Println("Erreur: Tous les champs (nom, prenom, tel) sont requis pour ajouter un contact")
			os.Exit(1)
		}

		if annuaire.ContactExists(*nom) {
			fmt.Printf("Erreur: Un contact avec le nom '%s' existe déjà\n", *nom)
			os.Exit(1)
		}

		annuaire.SetAnnuaire(*nom, *prenom, *tel)

	case "rechercher":
		if *nom == "" {
			fmt.Println("Erreur: Le nom est requis pour rechercher un contact")
			os.Exit(1)
		}
		annuaire.SearchAnnuaire(*nom)

	case "lister":
		annuaire.GetAnnuaire()

	case "supprimer":
		if *nom == "" {
			fmt.Println("Erreur: Le nom est requis pour supprimer un contact")
			os.Exit(1)
		}
		annuaire.DeleteAnnuaire(*nom)

	case "modifier":
		if *nom == "" {
			fmt.Println("Erreur: Le nom est requis pour modifier un contact")
			os.Exit(1)
		}
		annuaire.UpdateAnnuaire(*nom, *prenom, *tel)

	default:
		fmt.Println("Actions disponibles:")
		fmt.Println("  --action ajouter --nom \"Nom\" --prenom \"Prenom\" --tel \"0123456789\"")
		fmt.Println("  --action rechercher --nom \"Nom\"")
		fmt.Println("  --action lister")
		fmt.Println("  --action supprimer --nom \"Nom\"")
		fmt.Println("  --action modifier --nom \"Nom\" [--prenom \"NouveauPrenom\"] [--tel \"0123456789\"]")
		os.Exit(1)
	}
}
