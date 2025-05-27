package annuaire

import "fmt"

type Annuaire struct {
	nom    string
	prenom string
	numero string
}

var entries = []Annuaire{
	{nom: "emmanuel", prenom: "John", numero: "1234567890"},
	{nom: "laurent", prenom: "Jane", numero: "0987654321"},
}

func GetAnnuaire() {
	fmt.Println("Liste de l'annuaire:")
	fmt.Println("====================")

	if len(entries) == 0 {
		fmt.Println("Aucun contact dans l'annuaire")
		return
	}

	for i, entry := range entries {
		fmt.Printf("%d. Nom: %s, Prénom: %s, Numéro: %s\n",
			i+1, entry.nom, entry.prenom, entry.numero)
	}
}

// SetAnnuaire ajoute un nouveau contact à l'annuaire
func SetAnnuaire(nom, prenom, numero string) {
	fmt.Printf("Ajout du contact: %s %s (%s)\n", prenom, nom, numero)
	entries = append(entries, Annuaire{nom: nom, prenom: prenom, numero: numero})
	fmt.Println("Contact ajouté avec succès!")
}

// supprime un contact par nom
func DeleteAnnuaire(nom string) {
	fmt.Printf("Recherche du contact à supprimer: %s\n", nom)

	for i, entry := range entries {
		if entry.nom == nom {
			fmt.Printf("Suppression du contact: %s %s (%s)\n",
				entry.prenom, entry.nom, entry.numero)
			entries = append(entries[:i], entries[i+1:]...)
			fmt.Println("Contact supprimé avec succès!")
			return
		}
	}

	fmt.Printf("Aucun contact trouvé avec le nom: %s\n", nom)
}

// supprime tous les contacts
func DeleteAllAnnuaire() {
	fmt.Println("Suppression de tous les contacts...")

	count := len(entries)
	for i := len(entries) - 1; i >= 0; i-- {
		fmt.Printf("Suppression: %s %s (%s)\n",
			entries[i].prenom, entries[i].nom, entries[i].numero)
		entries = entries[:i]
	}

	fmt.Printf("%d contact(s) supprimé(s)\n", count)
}

func UpdateAnnuaire(nom, nouveauPrenom, nouveauNumero string) {
	fmt.Printf("Recherche du contact à modifier: %s\n", nom)

	for i, entry := range entries {
		if entry.nom == nom {
			fmt.Printf("Contact trouvé: %s %s (%s)\n",
				entry.prenom, entry.nom, entry.numero)

			if nouveauPrenom != "" {
				entries[i].prenom = nouveauPrenom
			}
			if nouveauNumero != "" {
				entries[i].numero = nouveauNumero
			}

			fmt.Printf("Contact modifié: %s %s (%s)\n",
				entries[i].prenom, entries[i].nom, entries[i].numero)
			fmt.Println("Contact mis à jour avec succès!")
			return
		}
	}

	fmt.Printf("Aucun contact trouvé avec le nom: %s\n", nom)
}

func SearchAnnuaire(nom string) {
	fmt.Printf("Recherche dans l'annuaire: %s\n", nom)
	fmt.Println("==============================")

	found := false
	for _, entry := range entries {
		if entry.nom == nom {
			fmt.Printf("Contact trouvé:\n")
			fmt.Printf("  Nom: %s\n", entry.nom)
			fmt.Printf("  Prénom: %s\n", entry.prenom)
			fmt.Printf("  Numéro: %s\n", entry.numero)
			found = true
		}
	}

	if !found {
		fmt.Printf("Aucun contact trouvé avec le nom: %s\n", nom)
	}
}

func ContactExists(nom string) bool {
	for _, entry := range entries {
		if entry.nom == nom {
			return true
		}
	}
	return false
}
