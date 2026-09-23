package main

import (
	"fmt"
)

// ==========================================
// MENU PRINCIPAL
// ==========================================

// Affiche le menu principal et redirige vers les sous-systèmes selon le choix du joueur
func MainMenu(c *Character) {
	for {
		fmt.Println("\n=== MENU PRINCIPAL ===")
		fmt.Println("1 - Afficher les infos du personnage")
		fmt.Println("2 - Afficher l'inventaire")
		fmt.Println("3 - Le marchand râleur")
		fmt.Println("4 - Le forgeron vantard")
		fmt.Println("5 - Qui sont-ils ?")
		fmt.Println("6 - Entraînement")
		fmt.Println("7 - Quitter")
		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			DisplayInfo(c)
		case 2:
			AccessInventory(c)
		case 3:
			Merchant(c)
		case 4:
			Blacksmith(c)
		case 5:
			WhoAreThey(c)
		case 6:
			trainingFight(c)
		case 7:
			fmt.Println("Allez ouste !")
			return
		default:
			fmt.Println("Choisis ce qui est proposé !")
		}
	}
}
