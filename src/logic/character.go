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
}

func main() {
	noah := initCharacter("Noah", "vagabond", 1, 100, 80, []string{"Katana Laser", "Revolver XRAY", "Fusil nucleaire"})

	fmt.Println("Nom :", noah.Name)
	fmt.Println("Classe :", noah.Classe)
	fmt.Println("Niveau :", noah.Lvl)
	fmt.Println("PV :", noah.CurrentHP, "/", noah.MaxHP)
	fmt.Println("Inventaire :", noah.Inventory)
}
