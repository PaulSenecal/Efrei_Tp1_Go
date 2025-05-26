package annuaire

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
	println("Fetching the annuaire...")

	for _, entry := range entries {
		println("Nom:", entry.nom, "Prenom:", entry.prenom, "Numero:", entry.numero)
	}
}
func SetAnnuaire(nom, prenom, numero string) {
	println("New entry added:", nom, prenom, numero)
	entries = append(entries, Annuaire{nom: nom, prenom: prenom, numero: numero})
}

func DeleteAnnuaire(nom string) {
	println("Deleting entry with nom:", nom)
	for i, entry := range entries {
		println("Deleting entry:", entry.nom, entry.prenom, entry.numero)
		entries = append(entries[:i], entries[i+1:]...)
		break
	}
}
func DeleteAllAnnuaire() {
	for i := 0; i < len(entries); i++ {
		println("Deleting entry:", entries[i].nom, entries[i].prenom, entries[i].numero)
		entries = append(entries[:i], entries[i+1:]...)
		i--
	}
}
func UpdateAnnuaire() {
	println("Update function not implemented yet.")
}
func SearchAnnuaire(nom string) {
	println("Searching in the annuaire...")
	for _, entry := range entries {
		if entry.nom == nom {
			println("Found:", entry.nom, entry.prenom, entry.numero)
			return
		}
	}
}
