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
}

func initCharacter(name, classe string, lvl, maxHP, currentHP int, inventory []string, sorts []string) Character {
	return Character{
		Name:      name,
		Classe:    classe,
		Lvl:       lvl,
		MaxHP:     maxHP,
		CurrentHP: currentHP,
		Inventory: inventory,
		Sorts:     sorts,
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
<<<<<<< HEAD
	fmt.Println("Sorts :", c.Sorts)
=======
	fmt.Println("Tapes 0 pour revenir au menu principal")
	var back int
	fmt.Scan(&back)
>>>>>>> eva
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
=======
	fmt.Println("Tapes 0 pour revenir au menu principal")
	var back int
	fmt.Scan(&back)


func MainMenu(c Character) {
	for {
		fmt.Println("\n=== MENU PRINCIPAL ===")
		fmt.Println("1 - Afficher les infos du personnage")
		fmt.Println("2 - Afficher l'inventaire")
		fmt.Println("3 - Quitter")
		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scan(&choice) // Récupère le nombre tapé par l'utilisateur (1, 2 ou 3)

		switch choice {
		case 1:
			DisplayInfo(c) // On exécute la fonction, puis la boucle 'for' reprend
		case 2:
			AccessInventory(c) // Idem, puis la boucle reprend
		case 3:
			fmt.Println("Aller ouste !")
			return // Seul le choix 3 utilise 'return' pour stopper la boucle et quitter le jeu
		default:
			fmt.Println("Choisis ce qui est proposé quiquiche")
		}
	}
}

func AddInventory(c *Character, item string) {
    c.Inventory = append(c.Inventory, item)
}

func Merchant(c *Character){
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
