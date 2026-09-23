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
		if !unicode.IsLetter(r) {
			return "", false
		}

		if i == 0 {
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
	for {
		fmt.Print("Entrez votre nom : ")
		fmt.Scan(&rawName)

		formatted, valid := processName(rawName)
		if !valid {
			fmt.Println("Erreur : Les numéros ne sont pas acceptés dans le pseudo !")
			continue
		}

		name = formatted
		break
	}

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
			fmt.Scanln(&dump)
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
			fmt.Println("Erreur : Choix invalide ! Veuillez saisir 1, 2 ou 3, pfff.")
			continue
		}
		break
	}

	player := initCharacter(name, classeName, 1, 100, 100, []string{"Potion de mana"}, []string{}, 100, 100)
	skill(&player)
	MainMenu(&player)
}
