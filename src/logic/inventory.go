package main

import (
	"fmt"
)

// ==========================================
// GESTION DE L'INVENTAIRE & ÉQUIPEMENT
// ==========================================

// Affiche l'inventaire du joueur et permet de sélectionner un objet
func AccessInventory(c *Character) {
	fmt.Println("\n=================================")
	fmt.Println("    INVENTAIRE DU PERSONNAGE    ")
	fmt.Println("=================================")
	if len(c.Inventory) == 0 {
		fmt.Println("Votre inventaire est vide")
		fmt.Println("Tapez 0 pour revenir au menu principal")
		var back int
		fmt.Scan(&back)
		return
	}

	for i, item := range c.Inventory {
		fmt.Println(i+1, "-", item)
	}

	fmt.Println("Choisissez un objet à utiliser (ou 0 pour quitter) :")
	var choice int
	fmt.Scan(&choice)

	if choice > 0 && choice <= len(c.Inventory) {
		item := c.Inventory[choice-1]
		useItem(c, item, choice-1)
	}
}

// Ajoute un objet à l'inventaire en vérifiant la capacité maximale
func addItem(c *Character, item string) bool {
	if len(c.Inventory) >= c.MaxInventory {
		fmt.Println("L'inventaire est plein ! Impossible d'ajouter :", item)
		return false
	}
	c.Inventory = append(c.Inventory, item)
	return true
}

// Retire une liste d'objets spécifiques de l'inventaire (ex: composants de forge)
func removeItems(c *Character, itemsNeeded []string) bool {
	tempInventory := append([]string{}, c.Inventory...)

	for _, needed := range itemsNeeded {
		found := false
		for i, item := range tempInventory {
			if item == needed {
				tempInventory = append(tempInventory[:i], tempInventory[i+1:]...)
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	c.Inventory = tempInventory
	return true
}

// Augmente la capacité maximale de l'inventaire
func upgradeInventorySlot(c *Character) {
	if c.InventoryUpgrades < 3 {
		c.MaxInventory += 10
		c.InventoryUpgrades += 1
		fmt.Println("Inventaire augmenté ! Capacité maximale :", c.MaxInventory)
	} else {
		fmt.Println("Inventaire au max. Arrête de cliquer.")
	}
}

// Apprend le sort Boule de Feu si le joueur ne le possède pas
func spellBook(c *Character) {
	for _, j := range c.Sorts {
		if j == "Boule de Feu" {
			fmt.Println("Vous connaissez déjà ce sort !")
			return
		}
	}
	c.Sorts = append(c.Sorts, "Boule de Feu")
	fmt.Println("Vous avez appris : Boule de Feu !")
}

// Équipe une pièce d'armure et applique les bonus de PV
func addEquipment(char *Character, stuff string) {
	for i, obj := range char.Inventory {
		if obj == stuff {
			char.Inventory = append(char.Inventory[:i], char.Inventory[i+1:]...)
			break
		}
	}

	switch stuff {
	case "Chapeau de l'aventurier":
		if char.Equipment.Tete != "" {
			char.Inventory = append(char.Inventory, char.Equipment.Tete)
			char.MaxHP -= 10
		}
		char.Equipment.Tete = stuff
		char.MaxHP += 10

	case "Tunique de l'aventurier":
		if char.Equipment.Torse != "" {
			char.Inventory = append(char.Inventory, char.Equipment.Torse)
			char.MaxHP -= 25
		}
		char.Equipment.Torse = stuff
		char.MaxHP += 25

	case "Bottes de l'aventurier":
		if char.Equipment.Pieds != "" {
			char.Inventory = append(char.Inventory, char.Equipment.Pieds)
			char.MaxHP -= 15
		}
		char.Equipment.Pieds = stuff
		char.MaxHP += 15

	default:
		fmt.Println("Cet équipement n'existe pas :", stuff)
	}
}

// Utilise un objet (consommable, équipement ou livre)
func useItem(c *Character, item string, index int) {
	if item == "Potion de vie" {
		if c.CurrentHP >= c.MaxHP {
			fmt.Println("Vos points de vie sont déjà au maximum !")
			return
		}
		c.CurrentHP += 50
		if c.CurrentHP > c.MaxHP {
			c.CurrentHP = c.MaxHP
		}
		fmt.Printf("Vous utilisez une Potion de Vie. PV actuels : %d/%d\n", c.CurrentHP, c.MaxHP)
		c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)

	} else if item == "Potion de poison" {
		fmt.Println("Vous buvez la potion de poison... T'as pas inventé l'eau chaude toi !")
		c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)
		poisonPot(c)

	} else if item == "Potion de mana" {
		if c.CurrentMana >= c.MaxMana {
			fmt.Println("Votre mana est déjà au maximum !")
			return
		}
		c.CurrentMana += 30
		if c.CurrentMana > c.MaxMana {
			c.CurrentMana = c.MaxMana
		}
		fmt.Printf("Vous utilisez une Potion de Mana. Mana actuel : %d/%d\n", c.CurrentMana, c.MaxMana)
		c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)

	} else if item == "Chapeau de l'aventurier" || item == "Tunique de l'aventurier" || item == "Bottes de l'aventurier" {
		addEquipment(c, item)
		fmt.Printf("Vous avez équipé : %s !\n", item)

	} else if item == "Livre de Sort : Boule de Feu" {
		spellBook(c)
		c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)

	} else {
		fmt.Println("Cet objet ne peut pas être consommé ou équipé directement.")
	}
}
