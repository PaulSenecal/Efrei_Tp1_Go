# Annuaire Go

Un programme simple de gestion d'annuaire téléphonique écrit en Go.

## Auteur
Paul Senecal - 2023-10-01

## Description
Ce programme permet de gérer un annuaire téléphonique avec les fonctionnalités de base : ajout, recherche, suppression et affichage des entrées.

## Structure du projet
```
tp1/
├── main.go          # Point d'entrée du programme
├── annuaire/
│   └── annuaire.go  # Package contenant la logique de l'annuaire
└── README.md        # Ce fichier
```

## Installation
1. Clonez le repository
2. Assurez-vous d'avoir Go installé sur votre système
3. Naviguez dans le répertoire du projet

## Utilisation
### Commande de base
```bash
go run main.go -nom "Nom" -prenom "Prenom" -telephone "NumeroTelephone"
```

### Exemples
```bash
# Ajouter une entrée
go run main.go -nom "Doe" -prenom "John" -telephone "1234567890"

# Utiliser les valeurs par défaut
go run main.go
```

### Paramètres
- `-nom` : Le nom de famille à ajouter dans l'annuaire
- `-prenom` : Le prénom à ajouter dans l'annuaire  
- `-telephone` : Le numéro de téléphone à ajouter dans l'annuaire

## Fonctionnalités disponibles

### Package `annuaire`
Le package annuaire propose les fonctions suivantes :

#### `GetAnnuaire()`
Affiche tous les entries de l'annuaire.

#### `SetAnnuaire(nom, prenom, numero string)`
Ajoute une nouvelle entrée à l'annuaire.

#### `SearchAnnuaire(nom string)`
Recherche une entrée par nom dans l'annuaire.

#### `DeleteAnnuaire(nom string)`
Supprime une entrée spécifique de l'annuaire basée sur le nom.

#### `DeleteAllAnnuaire()`
Supprime toutes les entrées de l'annuaire.



## Structure des données
```go
type Annuaire struct {
    nom    string  // Nom de famille
    prenom string  // Prénom
    numero string  // Numéro de téléphone
}
```

## Données par défaut
L'annuaire contient initialement deux entrées :
- emmanuel John - 1234567890
- laurent Jane - 0987654321

## Compilation
```bash
# Compiler le programme
go build main.go

# Exécuter le binaire compilé
./main -nom "Doe" -prenom "John" -telephone "1234567890"
```

## Exemples d'utilisation avancée
Pour utiliser toutes les fonctionnalités, vous pouvez décommenter et modifier la section dans `main.go` :

```go
annuaire.SetAnnuaire("Doe", "John", "1234567890")
annuaire.SetAnnuaire("Smith", "Jane", "0987654321")
annuaire.GetAnnuaire()
annuaire.SearchAnnuaire("Doe")
annuaire.DeleteAllAnnuaire()
annuaire.GetAnnuaire()
```

## Notes techniques
- Le programme utilise le package `flag` pour la gestion des arguments en ligne de commande
- Les données sont stockées en mémoire et ne persistent pas entre les exécutions
- La fonction de suppression individuelle a un comportement qui pourrait nécessiter une révision

## Améliorations possibles
- Implémenter la fonction `UpdateAnnuaire()`
- Ajouter la persistance des données (fichier, base de données)
- Améliorer la gestion des erreurs
- Ajouter des validations pour les numéros de téléphone
- Implémenter une interface utilisateur plus interactive

## Licence
Ce projet est à des fins éducatives.