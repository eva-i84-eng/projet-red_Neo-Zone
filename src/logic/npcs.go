package main

import (
	"fmt"
)

// ==========================================
// MARCHANDS & PNJS
// ==========================================

// Boutique du marchand râleur : achat d'objets, potions, livres et améliorations
func Merchant(c *Character) {
	for {
		fmt.Println("\n=== MARCHAND RÂLEUR ===")
		fmt.Println("Or disponible :", c.Money)
		fmt.Println("1 - Potion de vie (3 pièces d'or)")
		fmt.Println("2 - Potion de poison (6 pièces d'or)")
		fmt.Println("3 - Livre de Sort : Boule de Feu (25 pièces d'or)")
		fmt.Println("4 - Fourrure de Loup (4 pièces d'or)")
		fmt.Println("5 - Peau de Troll (7 pièces d'or)")
		fmt.Println("6 - Cuir de Sanglier (3 pièces d'or)")
		fmt.Println("7 - Plume de Corbeau (1 pièce d'or)")
		fmt.Println("8 - Augmenter l'inventaire (30 pièces d'or)")
		fmt.Println("9 - Potion de mana (5 pièces d'or)")
		fmt.Println("0 - Retour au menu principal")
		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scan(&choice)

		var item string
		var price int

		switch choice {
		case 1:
			item = "Potion de vie"
			price = 3
		case 2:
			item = "Potion de poison"
			price = 6
		case 3:
			item = "Livre de Sort : Boule de Feu"
			price = 25
		case 4:
			item = "Fourrure de Loup"
			price = 4
		case 5:
			item = "Peau de Troll"
			price = 7
		case 6:
			item = "Cuir de Sanglier"
			price = 3
		case 7:
			item = "Plume de Corbeau"
			price = 1
		case 8:
			if c.Money >= 30 {
				c.Money -= 30
				upgradeInventorySlot(c)
			} else {
				fmt.Println("T'es fauché. Upgrade refusé.")
			}
			continue
		case 9:
			item = "Potion de mana"
			price = 5
		case 0:
			fmt.Println("C'est bon j'ai compris, dégage !")
			return
		default:
			fmt.Println("Choix invalide !")
			continue
		}

		if c.Money >= price {
			if addItem(c, item) {
				c.Money -= price
				fmt.Printf("Tu as acheté : %s pour %d pièces d'or.\n", item, price)
			}
		} else {
			fmt.Println("T'as pas assez de thunes !")
		}
	}
}

// Atelier du forgeron : fabrication d'équipements à partir de composants
func Blacksmith(c *Character) {
	for {
		fmt.Println("\n=== LE FORGERON ===")
		fmt.Println("Or disponible :", c.Money)
		fmt.Println("1 - Chapeau de l'aventurier (1 Plume de Corbeau, 1 Cuir de Sanglier)")
		fmt.Println("2 - Tunique de l'aventurier (2 Fourrures de Loup, 1 Peau de Troll)")
		fmt.Println("3 - Bottes de l'aventurier  (1 Fourrure de Loup, 1 Cuir de Sanglier)")
		fmt.Println("0 - Retour au menu principal")
		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scan(&choice)

		var equipment string
		var neededItems []string
		cost := 5

		switch choice {
		case 1:
			equipment = "Chapeau de l'aventurier"
			neededItems = []string{"Plume de Corbeau", "Cuir de Sanglier"}
		case 2:
			equipment = "Tunique de l'aventurier"
			neededItems = []string{"Fourrure de Loup", "Fourrure de Loup", "Peau de Troll"}
		case 3:
			equipment = "Bottes de l'aventurier"
			neededItems = []string{"Fourrure de Loup", "Cuir de Sanglier"}
		case 0:
			return
		default:
			fmt.Println("Choix invalide !")
			continue
		}

		if c.Money < cost {
			fmt.Println("Pas assez de pièces !")
			continue
		}

		if !removeItems(c, neededItems) {
			fmt.Println("Il te manque des composants !")
			continue
		}

		c.Money -= cost
		addItem(c, equipment)
		fmt.Printf("Fabriqué : %s pour %d pièces d'or !\n", equipment, cost)
	}
}

// Rencontres spéciales avec les artistes cachés
func WhoAreThey(c *Character) {
	for {
		fmt.Println("\n=== LES ARTISTES CACHÉS ===")
		fmt.Println("1 - Steven Spielberg")
		fmt.Println("2 - ABBA")
		fmt.Println("0 - Retour")
		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			if c.HasTalkedToSpielberg {
				fmt.Println("Vous avez déjà salué S. Spielberg !")
			} else {
				fmt.Println("Vous avez salué S.Spielberg ! Il vous raconte une histoire et vous offre 10 pièces d'or !")
				c.Money += 10
				c.HasTalkedToSpielberg = true
			}
		case 2:
			if c.HasTalkedToABBA {
				fmt.Println("Vous avez déjà écouté ABBA !")
			} else {
				fmt.Println("ABBA vous chante une chanson, vos PV sont restaurés !")
				c.CurrentHP = c.MaxHP
				c.HasTalkedToABBA = true
			}
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
