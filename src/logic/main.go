package main

import (
	"fmt"
	"unicode"
)

func processName(input string) (string, bool) {
	if len(input) == 0 {
		return "", false
	}

	runes := []rune(input)

	for i, r := range runes {
		// On autorise lettres ET espaces
		if !unicode.IsLetter(r) && !unicode.IsSpace(r) {
			return "", false
		}

		if i == 0 || (i > 0 && unicode.IsSpace(runes[i-1])) {
			runes[i] = unicode.ToUpper(r)
		} else {
			runes[i] = unicode.ToLower(r)
		}
	}

	return string(runes), true
}

package main

import (
	"fmt"
	"unicode"
)

// Traitement du nom : autorise lettres et espaces,
// puis met une majuscule au début de chaque mot.
func processName(input string) (string, bool) {
	if len(input) == 0 {
		return "", false
	}

	runes := []rune(input)

	for i, r := range runes {
		// On autorise uniquement lettres et espaces
		if !unicode.IsLetter(r) && !unicode.IsSpace(r) {
			return "", false
		}

		// Majuscule au premier caractère ou juste après un espace
		if i == 0 || (i > 0 && unicode.IsSpace(runes[i-1])) {
			runes[i] = unicode.ToUpper(r)
		} else {
			runes[i] = unicode.ToLower(r)
		}
	}

	return string(runes), true
}

func main() {
	var rawName string
	var name string
	var classeChoice int
	var classeName string

	fmt.Println("=== CRÉATION DU PERSONNAGE ===")

	// Saisie du nom
	for {
		fmt.Print("Entrez votre nom : ")
		fmt.Scanln(&rawName)

		formatted, valid := processName(rawName)
		if !valid {
			fmt.Println("Erreur : Les chiffres et symboles ne sont pas acceptés dans le pseudo !")
			continue
		}

		name = formatted
		break
	}

	// Choix de la classe
	for {
		fmt.Println("\nChoisissez votre classe :")
		fmt.Println("1 - Vagabond")
		fmt.Println("2 - Samurai")
		fmt.Println("3 - Cowboy")
		fmt.Print("Votre choix : ")

		_, err := fmt.Scan(&classeChoice)
		if err != nil {
			fmt.Println("Erreur : Veuillez entrer un NOMBRE valide.")
			var dump string
			fmt.Scanln(&dump) // Vide le choix invalide
			continue
		}

		switch classeChoice {
		case 1:
			classeName = "Vagabond"
		case 2:
			classeName = "Samurai"
		case 3:
			classeName = "Cowboy"
		default:
			fmt.Println("Erreur : Choix invalide ! Veuillez saisir 1, 2 ou 3.")
			continue
		}
		break
	}

	// Initialisation et lancement du jeu
	player := initCharacter(name, classeName, 1, 100, 100, []string{"Potion de mana"}, []string{}, 100, 100)
	skill(&player)
	MainMenu(&player)
}
