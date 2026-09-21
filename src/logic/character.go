package main

import "fmt"

type Character struct {
	Name       string
	Classe     string
	Lvl        int
	MaxHP      int
	CurrentHP  int
	MaxHPFinal int
	Inventory  []string
	Money      int
	Sorts      []string
}

func initCharacter(name, classe string, lvl, maxHP, currentHP int, inventory []string, sorts []string) Character {
	return Character{
		Name:       name,
		Classe:     classe,
		Lvl:        lvl,
		MaxHP:      maxHP,
		CurrentHP:  currentHP,
		MaxHPFinal: 600,
		Inventory:  inventory,
		Sorts:      sorts,
		Money:      100,
	}
}

func main() {
	noah := initCharacter("Noah", "vagabond", 1, 100, 80, []string{"Katana Laser", "Revolver XRAY"}, []string{"Coup de Poing"})
	MainMenu(&noah)
}

func DisplayInfo(c *Character) {
	fmt.Println("Nom :", c.Name)
	fmt.Println("Classe :", c.Classe)
	fmt.Println("Niveau :", c.Lvl)
	fmt.Println("PV :", c.CurrentHP, "/", c.MaxHP)
	fmt.Println("Or :", c.Money)
	fmt.Println("Inventaire :", c.Inventory)
	fmt.Println("Tapes 0 pour revenir au menu principal")
	var back int
	fmt.Scan(&back)
}

func characterCreation() Character {
	var name string
	valid := false

	for !valid {
		fmt.Print("Entrez le nom de votre personnage : ")
		fmt.Scan(&name)

		valid = true
		for _, c := range name {
			if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') {
				valid = false
			}
		}

		if !valid {
			fmt.Println("Nom invalide, uniquement des lettres.")
		}
	}

	result := ""
	for i, c := range name {
		if i == 0 && c >= 'a' && c <= 'z' {
			c = c - 32
		}
		if i > 0 && c >= 'A' && c <= 'Z' {
			c = c + 32
		}
		result += string(c)
	}

	var classe string
	var maxHP int
	var inventory []string

	for {
		fmt.Print("Choisissez votre classe (Samurai, Cowboy) : ")
		fmt.Scan(&classe)

		if classe == "Samurai" {
			maxHP = 80
			inventory = []string{"katana laser", "shuriken G300", "Odachi"}
			break
		} else if classe == "Cowboy" {
			maxHP = 100
			inventory = []string{"Revolver XRAY", "Fusil nucléaire", "Fusil Evans"}
			break
		}
		fmt.Println("Votre classe n'est pas disponible. Choisissez Samurai ou Cowboy.")
	}

	currentHP := maxHP / 2

	return initCharacter(result, classe, 1, maxHP, currentHP, inventory, []string{"Coup de Poing"})
}

func AccessInventory(c *Character) {
	fmt.Println("=================================")
	fmt.Println("     INVENTAIRE DU PERSONNAGE    ")
	fmt.Println("=================================")
	if len(c.Inventory) <= 0 {
		fmt.Println("Votre inventaire est vide")
	} else {
		for i, item := range c.Inventory {
			fmt.Println(i+1, "-", item)
		}
	}

	fmt.Println("Tapes 0 pour revenir au menu principal")
	var back int
	fmt.Scan(&back)
}

func MainMenu(c *Character) {
	for {
		fmt.Println("\n=== MENU PRINCIPAL ===")
		fmt.Println("1 - Afficher les infos du personnage")
		fmt.Println("2 - Afficher l'inventaire")
		fmt.Println("3 - Le marchand raleur")
		fmt.Println("4 - Le forgeron vantard")
		fmt.Println("5 - Quitter")
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
			fmt.Println("Aller ouste !")
			return
		default:
			fmt.Println("Choisis ce qui est proposé quiquiche")
		}
	}
}

func AddInventory(c *Character, item string) {
	c.Inventory = append(c.Inventory, item)
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
			fmt.Println("C'est bon j'ai compris, dégages !")
			return
		default:
			fmt.Println("Choix invalide, achète quelque chose de vrai !")
			continue
		}

		if c.Money >= price {
			c.Money -= price
			AddInventory(c, item)
			fmt.Printf("Tu as acheté : %s pour %d pièces d'or.\n", item, price)
		} else {
			fmt.Println("T'as pas assez de thunes, repasses quand tu seras riche !")
		}
	}
}

func Blacksmith(c *Character) {
	for {
		fmt.Println("\n=== LE FORGERON ===")
		fmt.Println("Or disponible :", c.Money)
		fmt.Println("Fabriquer un équipement coûte 5 pièces d'or.")
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
			fmt.Println("Tu reviendras vite de toute façon, je suis le meilleur forgeron !")
			return
		default:
			fmt.Println("Je pourrais le faire hein, je peux TOUT faire mais j'ai pas envie !")
			continue
		}
		if c.Money < cost {
			fmt.Println("Pas assez de pièces ! Tu veux une réduc ? Bah nan, je suis déjà trop gentil comme ça !")
			continue
		}
		if !removeItems(c, neededItems) {
			fmt.Println("Il te manque des composants ! Quoi, tu croyais que j'allais faire de la magie sans matières premières ? Repasse quand tu auras le matos !")
			continue
		}
		c.Money -= cost
		AddInventory(c, equipment)
		fmt.Printf("Regarde-moi cette merveille ! Je t'ai fabriqué : %s pour %d pièces d'or !\n", equipment, cost)
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

func addItem(c *Character, item string) {
	if len(c.Inventory) >= 10 {
		fmt.Println("L'inventaire est plein. Nous ne pouvons accepter :", item)
		return
	}
	c.Inventory = append(c.Inventory, item)
}


func skill(c *Character) {
	if c.Classe == "Samurai" {
		c.Sorts = []string{"Tempete du ninja", "Coup de poing"}
		return
	} else if c.Classe == "Cowboy" {
		c.Sorts = []string{"Slowing Time", "Coup de poing"}
		return
	}
	c.Sorts = []string{"Coup de poing"}
}

func spellBook(c *Character) {
	for _, j := range c.Sorts {
		if j == "Boule de Feu" {
			return
		}
	}
	c.Sorts = append(c.Sorts, "Boule de Feu")
}

type Equipment struct {
	tete  string
	torse string
	pieds string
}
>>>>>>> 4ed7e7b08df46ba947a7abce41dae5aaa5221713
