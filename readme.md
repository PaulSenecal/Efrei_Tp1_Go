# TP1-Annuaire

## Description
Programme d'annuaire simple en Go permettant de gérer des contacts avec nom, prénom et numéro de téléphone.

## Membres du groupe
- Paul Senecal

## Structure du projet
```
tp1/
├── main.go           # Point d'entrée du programme avec gestion des flags
├── annuaire/
│   ├── annuaire.go   # Package contenant la logique métier
│   └── annuaire_test.go # Tests unitaires
├── go.mod            # Module Go
└── README.md         # Documentation
```

## Fonctionnalités

### Actions disponibles
- **Ajouter** un contact
- **Rechercher** un contact par nom
- **Lister** tous les contacts
- **Supprimer** un contact
- **Modifier** un contact existant

### Fonctionnalités avancées
- Vérification de l'existence d'un contact avant ajout (évite les doublons)
- Modification partielle des contacts (prénom et/ou téléphone)
- Messages d'erreur informatifs
- Interface en ligne de commande intuitive

## Installation et utilisation

### Prérequis
- Go 1.19 ou plus récent

### Installation
```bash
git clone <lien-du-repo>
cd tp1
go mod init github.com/PaulSenecal/tp1
```

### Commandes d'utilisation

#### Ajouter un contact
```bash
go run main.go --action ajouter --nom "Charlie" --prenom "John" --tel "0811223344"
```

#### Rechercher un contact
```bash
go run main.go --action rechercher --nom "Alice"
```

#### Lister tous les contacts
```bash
go run main.go --action lister
```

#### Supprimer un contact
```bash
go run main.go --action supprimer --nom "Charlie"
```

#### Modifier un contact
```bash
# Modifier le prénom et le téléphone
go run main.go --action modifier --nom "Charlie" --prenom "Charles" --tel "0123456789"

# Modifier seulement le prénom
go run main.go --action modifier --nom "Charlie" --prenom "Charles"

# Modifier seulement le téléphone
go run main.go --action modifier --nom "Charlie" --tel "0123456789"
```

#### Aide
```bash
go run main.go
# Affiche la liste des actions disponibles
```

## Tests unitaires

### Exécuter les tests
```bash
cd annuaire
go test -v
```

### Tests implémentés
1. **TestSetAnnuaire** - Teste l'ajout de contacts
2. **TestContactExists** - Teste la vérification d'existence de contacts
3. **TestDeleteAnnuaire** - Teste la suppression de contacts
4. **TestUpdateAnnuaire** - Teste la modification de contacts

## Exemples d'utilisation

### Scénario complet
```bash
# Lister l'annuaire initial
go run main.go --action lister

# Ajouter un nouveau contact
go run main.go --action ajouter --nom "Dupont" --prenom "Marie" --tel "0606060606"

# Rechercher le contact ajouté
go run main.go --action rechercher --nom "Dupont"

# Modifier le contact
go run main.go --action modifier --nom "Dupont" --prenom "Marie-Claire" --tel "0707070707"

# Supprimer le contact
go run main.go --action supprimer --nom "Dupont"

# Vérifier la suppression
go run main.go --action lister
```

## Gestion d'erreurs
- Vérification des paramètres requis pour chaque action
- Détection des contacts déjà existants lors de l'ajout
- Messages d'erreur explicites
- Codes de sortie appropriés

## Architecture
Le projet utilise une architecture modulaire avec :
- **main.go** : Interface utilisateur et gestion des flags
- **annuaire/annuaire.go** : Logique métier et manipulation des données
- **annuaire/annuaire_test.go** : Tests unitaires complets

Les données sont stockées en mémoire dans un slice de structures `Annuaire`.