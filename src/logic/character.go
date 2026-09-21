package main

import "fmt"

type Character struct {
	Name      string
	Classe    string
	Lvl       int
	MaxHP     int
	CurrentHP int
	Inventory []string
	Sorts     []string
	Equipment Equipment
}

func initCharacter(name, classe string, lvl, maxHP, currentHP int, inventory []string, sorts, equipment Equipment) Character {
	return Character{
		Name:      name,
		Classe:    classe,
		Lvl:       lvl,
		MaxHP:     maxHP,
		CurrentHP: currentHP,
		Inventory: inventory,
		Sorts:     sorts,
		Equipment: equipment,
	}
}

func main() {
	// Création du personnage via la fonction
	noah := initCharacter("Noah", "vagabond", 1, 100, 80, []string{"Katana Laser", "Revolver XRAY", "Fusil Nucléaire"}, []string{"Coup de Poing"})

	// Affichage des informations
	DisplayInfo(noah)
}

func DisplayInfo(c Character) {
	fmt.Println("Nom :", c.Name)
	fmt.Println("Classe :", c.Classe)
	fmt.Println("Niveau :", c.Lvl)
	fmt.Println("PV :", c.CurrentHP, "/", c.MaxHP)
	fmt.Println("Inventaire :", c.Inventory)
	fmt.Println("Sorts :", c.Sorts)
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

func AccessInventory(c Character) {
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


func addEquipment(char *Character, stuff string) {
	if stuff == "Chapeau de l'aventurier" {
		if char.Equipment.Tete != "" {
			char.Inventory = append(char.Inventory, char.Equipment.Tete)
			char.MaxHP -= 10
		}
		char.Equipment.Tete = stuff
		char.MaxHP += 10
	} else if stuff == "Tunique de l'aventurier" {
		if char.Equipment.Torse != "" {
			char.Inventory = append(char.Inventory, char.Equipment.Torse)
			char.MaxHP -= 25
		}
		char.Equipment.Torse = stuff
		char.MaxHP += 25
	} else if stuff == "Bottes de l'aventurier" {
		if char.Equipment.Pieds != "" {
			char.Inventory = append(char.Inventory, char.Equipment.Pieds)
			char.MaxHP -= 15
		}
		char.Equipment.Pieds = stuff
		char.MaxHP += 15
	} else {
		fmt.Println("Cet équipement n'existe pas :", stuff)
		return
	}

	for i, obj := range char.Inventory {
		if obj == stuff {
			char.Inventory = append(char.Inventory[:i], char.Inventory[i+1:]...)
			break
		}
	}
}