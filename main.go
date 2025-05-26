package main

import (
	"flag"
	"fmt"

	"github.com/PaulSenecal/tp1/annuaire"
)

func main() {

	/*
		entries := []Annuaire{
			{nom: "Doe", prenom: "John", numero: "1234567890"},
			{nom: "Smith", prenom: "Jane", numero: "0987654321"},
		}*/

	var nom = "default text"
	flag.StringVar(&nom, "text", "default text", "Text to display")
	flag.Parse()

	/*var prenom = "default text"
	flag.StringVar(&prenom, "text", "default text", "Text to display")
	flag.Parse()*/

	fmt.Println("Programme Annuaire")

	annuaire.GetAnnuaire()
	annuaire.SetAnnuaire("Doe", "John", "1234567890")
	annuaire.SetAnnuaire("Smith", "Jane", "0987654321")
	annuaire.GetAnnuaire()
	annuaire.SearchAnnuaire("Doe")
	annuaire.DeleteAllAnnuaire()
	annuaire.GetAnnuaire()

}
