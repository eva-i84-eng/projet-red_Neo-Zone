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
