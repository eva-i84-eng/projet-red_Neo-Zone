package main

import "fmt"

type Equipment struct {
	Tete  string
	Torse string
	Pieds string
}

type Character struct {
	Name              string
	Classe            string
	Lvl               int
	MaxHP             int
	CurrentHP         int
	MaxHPFinal        int
	Inventory         []string
	MaxInventory      int
	InventoryUpgrades int
	Money             int
	Sorts             []string
	Equipment         Equipment
}

type Monster struct {
	Name      string
	MaxHP     int
	CurrentHP int
	Damage    int
}

func initCharacter(name, classe string, lvl, maxHP, currentHP int, inventory []string, sorts []string) Character {
	return Character{
		Name:              name,
		Classe:            classe,
		Lvl:               lvl,
		MaxHP:             maxHP,
		CurrentHP:         currentHP,
		MaxHPFinal:        600,
		Inventory:         inventory,
		MaxInventory:      10,
		InventoryUpgrades: 0,
		Sorts:             sorts,
		Money:             100,
	}
}

func main() {
	noah := initCharacter("Noah", "vagabond", 1, 100, 80, []string{"Katana Laser", "Revolver XRAY"}, []string{"Coup de Poing"})
	MainMenu(&noah)
}

func DisplayInfo(c *Character) {
	fmt.Println("\n=== INFORMATIONS ===")
	fmt.Println("Nom :", c.Name)
	fmt.Println("Classe :", c.Classe)
	fmt.Println("Niveau :", c.Lvl)
	fmt.Println("PV :", c.CurrentHP, "/", c.MaxHP)
	fmt.Println("Or :", c.Money)
	fmt.Println("Sorts :", c.Sorts)
	fmt.Println("Équipement :", c.Equipment)
	fmt.Println("Tapez 0 pour revenir au menu principal")
	var back int
	fmt.Scan(&back)
}

func AccessInventory(c *Character) {
	fmt.Println("\n=================================")
	fmt.Println("    INVENTAIRE DU PERSONNAGE    ")
	fmt.Println("=================================")
	if len(c.Inventory) == 0 {
		fmt.Println("Votre inventaire est vide")
	} else {
		for i, item := range c.Inventory {
			fmt.Println(i+1, "-", item)
		}
	}

	fmt.Println("Tapez 0 pour revenir au menu principal")
	var back int
	fmt.Scan(&back)
}

func addItem(c *Character, item string) bool {
	if len(c.Inventory) >= c.MaxInventory {
		fmt.Println("L'inventaire est plein ! Impossible d'ajouter :", item)
		return false
	}
	c.Inventory = append(c.Inventory, item)
	return true
}

func MainMenu(c *Character) {
	for {
		fmt.Println("\n=== MENU PRINCIPAL ===")
		fmt.Println("1 - Afficher les infos du personnage")
		fmt.Println("2 - Afficher l'inventaire")
		fmt.Println("3 - Le marchand râleur")
		fmt.Println("4 - Le forgeron vantard")
		fmt.Println("5 - Qui sont-ils ?")
		fmt.Println("6 - Quitter")
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
			fmt.Println("Aller ouste !")
			return
		default:
			fmt.Println("Choisis ce qui est proposé !")
		}
	}
}

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

func WhoAreThey(c *Character) {
	for {
		fmt.Println("\n=== LES ARTISTES CACHÉS ===")
		fmt.Println("1 - Diana")
		fmt.Println("2 - Michael")
		fmt.Println("0 - Retour")
		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Vous avez salué Diana ! Elle vous offre 10 pièces d'or.")
			c.Money += 10
			return
		case 2:
			fmt.Println("Michael vous chante une chanson, vos PV sont restaurés !")
			c.CurrentHP = c.MaxHP
			return
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}