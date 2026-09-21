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
		fmt.Println("4 - Quitter")
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
	fmt.Println("\n=== C'EST PAS CHER, PROMIS ===")
	fmt.Println("1 - Potion de vie (Gratuit)")
	fmt.Println("0 - Retour")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		AddInventory(c, "Potion de vie")
		fmt.Println("Aller c'est dans l'inventaire et utilise la bien t'as une sale tête")
	case 0:
		fmt.Println("C'est bon j'ai compris, dégages")
		return
	}
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
