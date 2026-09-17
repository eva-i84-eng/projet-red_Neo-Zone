package main

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

	_ = noah
}
