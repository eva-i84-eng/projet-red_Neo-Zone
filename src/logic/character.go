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

func main() {
	noah := Character{
		Name:      "Noah",
		Classe:    "vagabond",
		Lvl:       1,
		MaxHP:     100,
		CurrentHP: 80,
		Inventory: []string{"Katana Laser", "Revolver XRAY", "Fusil nucleaire"},
	}

	fmt.Println("Nom :", noah.Name)
	fmt.Println("Classe :", noah.Classe)
	fmt.Println("Niveau :", noah.Lvl)
	fmt.Println("PV :", noah.CurrentHP, "/", noah.MaxHP)
	fmt.Println("Inventaire :", noah.Inventory)
}
