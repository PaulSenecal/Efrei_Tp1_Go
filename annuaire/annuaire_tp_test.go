package annuaire

import (
	"testing"
)

func TestSetAnnuaire(t *testing.T) {
	originalEntries := make([]Annuaire, len(entries))
	copy(originalEntries, entries)

	initialCount := len(entries)
	SetAnnuaire("Testeur", "John", "0123456789")

	if len(entries) != initialCount+1 {
		t.Errorf("Attendu %d contacts, mais got %d", initialCount+1, len(entries))
	}

	lastEntry := entries[len(entries)-1]
	if lastEntry.nom != "Testeur" || lastEntry.prenom != "John" || lastEntry.numero != "0123456789" {
		t.Errorf("Contact ajouté incorrectement. Attendu: Testeur John 0123456789, Got: %s %s %s",
			lastEntry.nom, lastEntry.prenom, lastEntry.numero)
	}

	entries = originalEntries
}

func TestContactExists(t *testing.T) {
	if !ContactExists("emmanuel") {
		t.Error("ContactExists devrait retourner true pour 'emmanuel'")
	}

	if ContactExists("PersonneInexistante") {
		t.Error("ContactExists devrait retourner false pour 'PersonneInexistante'")
	}
}

func TestDeleteAnnuaire(t *testing.T) {
	originalEntries := make([]Annuaire, len(entries))
	copy(originalEntries, entries)

	SetAnnuaire("TempContact", "Test", "0000000000")
	initialCount := len(entries)

	DeleteAnnuaire("TempContact")

	if len(entries) != initialCount-1 {
		t.Errorf("Attendu %d contacts après suppression, mais got %d", initialCount-1, len(entries))
	}

	if ContactExists("TempContact") {
		t.Error("Le contact 'TempContact' devrait être supprimé")
	}

	entries = originalEntries
}

func TestUpdateAnnuaire(t *testing.T) {

	originalEntries := make([]Annuaire, len(entries))
	copy(originalEntries, entries)
	SetAnnuaire("TempUpdate", "Original", "1111111111")

	UpdateAnnuaire("TempUpdate", "Modified", "2222222222")

	found := false
	for _, entry := range entries {
		if entry.nom == "TempUpdate" {
			if entry.prenom != "Modified" || entry.numero != "2222222222" {
				t.Errorf("Contact mal modifié. Attendu: TempUpdate Modified 2222222222, Got: %s %s %s",
					entry.nom, entry.prenom, entry.numero)
			}
			found = true
			break
		}
	}

	if !found {
		t.Error("Contact 'TempUpdate' introuvable après modification")
	}

	entries = originalEntries
}
