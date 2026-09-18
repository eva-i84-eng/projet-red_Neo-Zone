package main

import "fmt"

type Character struct {
	Name      string
	Classe    string
	Lvl       int
	MaxHP     int
	CurrentHP int
	Inventory []string
}

func initCharacter(name, classe string, lvl, maxHP, currentHP int, inventory []string) Character {
	return Character{
		Name:      name,
		Classe:    classe,
		Lvl:       lvl,
		MaxHP:     maxHP,
		CurrentHP: currentHP,
		Inventory: inventory,
	}
	DisplayInfo(noah)
}

func DisplayInfo(c Character) {
	fmt.Println("Nom :", c.Name)
	fmt.Println("Classe :", c.Classe)
	fmt.Println("Niveau :", c.Lvl)
	fmt.Println("PV :", c.CurrentHP, "/", c.MaxHP)
	fmt.Println("Inventaire :", c.Inventory)
}