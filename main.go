/* Author: Paul Senecal
*	Date: 2023-10-01
*	Description: Programme d'annuaire
*	Usage: go run main.go -nom "Doe" -prenom "John" -telephone "1234567890"
*	Example: go run main.go -nom "Doe" -prenom "John" -telephone "1234567890"
 */

package main

import (
	"flag"
	"fmt"

	"github.com/PaulSenecal/tp1/annuaire"
)

func main() {

	var nom = "default text"
	flag.StringVar(&nom, "nom", "default nom", "le nom a ajouter dans l'annuaire")

	var prenom = "default text"
	flag.StringVar(&prenom, "prenom", "default prenom", "le prenom a ajouter dans l'annuaire")

	var telephone = "default text"
	flag.StringVar(&telephone, "telephone", "default telephone", "le tel a ajouter dans l'annuaire")
	flag.Parse()

	fmt.Println("Programme Annuaire")

	annuaire.SetAnnuaire(nom, prenom, telephone)

	/* utilisation des foncitons
	annuaire.SetAnnuaire("Doe", "John", "1234567890")
	annuaire.SetAnnuaire("Smith", "Jane", "0987654321")
	annuaire.GetAnnuaire()
	annuaire.SearchAnnuaire("Doe")
	annuaire.DeleteAllAnnuaire()
	annuaire.GetAnnuaire()*/
}
