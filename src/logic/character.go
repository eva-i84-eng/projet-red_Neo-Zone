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
	// Création du personnage via la fonction
	noah := initCharacter("Noah", "vagabond", 1, 100, 80, []string{"Katana Laser", "Revolver XRAY"})

	// Affichage des informations
	DisplayInfo(noah)
}

func DisplayInfo(c Character) {
	fmt.Println("Nom :", c.Name)
	fmt.Println("Classe :", c.Classe)
	fmt.Println("Niveau :", c.Lvl)
	fmt.Println("PV :", c.CurrentHP, "/", c.MaxHP)
	fmt.Println("Inventaire :", c.Inventory)
}