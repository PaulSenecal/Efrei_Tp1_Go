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
	/*println("Fetching the annuaire...")
	entries := []Annuaire{
		{nom: "emmanuel", prenom: "John", numero: "1234567890"},
		{nom: "laurent", prenom: "Jane", numero: "0987654321"},
	}*/

	for _, entry := range entries {
		println("Nom:", entry.nom, "Prenom:", entry.prenom, "Numero:", entry.numero)
	}
}
func SetAnnuaire(nom, prenom, numero string) {
	println("New entry added:", nom, prenom, numero)
}
func DeleteAnnuaire() {}
func UpdateAnnuaire() {}
func SearchAnnuaire(nom string) {
	println("Searching in the annuaire...")

}
