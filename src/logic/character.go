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
	DisplayInfo(noah)
}

func DisplayInfo(c Character) {
	fmt.Println("Nom :", c.Name)
	fmt.Println("Classe :", c.Classe)
	fmt.Println("Niveau :", c.Lvl)
	fmt.Println("PV :", c.CurrentHP, "/", c.MaxHP)
	fmt.Println("Inventaire :", c.Inventory)
}