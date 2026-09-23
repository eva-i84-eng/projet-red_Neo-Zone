package main

import (
	"fmt"
)

// ==========================================
// 1. STRUCTURES (MODÈLES DE DONNÉES)
// ==========================================

type Equipment struct {
	Tete  string
	Torse string
	Pieds string
}

type Character struct {
	Name                 string
	Classe               string
	Lvl                  int
	MaxHP                int
	CurrentHP            int
	MaxHPFinal           int
	Inventory            []string
	MaxInventory         int
	InventoryUpgrades    int
	Money                int
	Sorts                []string
	Equipment            Equipment
	MaxMana              int
	CurrentMana          int
	Initiative           int
	CurrentExp           int
	MaxExp               int
	BonusDamage          int // Dégâts bonus gagnés avec les niveaux
	HasTalkedToSpielberg bool
	HasTalkedToABBA      bool
}

type Monster struct {
	Name       string
	Lvl        int
	MaxHP      int
	CurrentHP  int
	Damage     int
	Initiative int
	Exp        int
}

// ==========================================
// 2. INITIALISATIONS DES ENTITÉS
// ==========================================

func initCharacter(name, classe string, lvl, maxHP, currentHP int, inventory []string, sorts []string, maxMana, currentMana int) Character {
	return Character{
		Name:              name,
		Classe:            classe,
		Lvl:               lvl,
		MaxHP:             maxHP,
		CurrentHP:         currentHP,
		MaxHPFinal:        600,
		Inventory:         inventory,
		MaxInventory:      10,
		InventoryUpgrades: 0,
		Sorts:             sorts,
		Money:             100,
		CurrentExp:        0,
		MaxExp:            100,
		BonusDamage:       0,
		MaxMana:           maxMana,
		CurrentMana:       currentMana,
		Initiative:        10,
	}
}

// Le gobelin s'adapte au niveau du joueur
func initGoblin(m *Monster, playerLvl int) {
	m.Lvl = playerLvl
	m.Name = fmt.Sprintf("Gobelin d'entraînement (Niv. %d)", m.Lvl)
	// Les statistiques du gobelin augmentent avec son niveau
	m.MaxHP = 40 + (playerLvl-1)*15
	m.CurrentHP = m.MaxHP
	m.Damage = 5 + (playerLvl-1)*3
	m.Initiative = 5
	m.Exp = 25 + (playerLvl-1)*10
}

// ==========================================
// 3. FONCTIONS LIÉES À L'ÉTAT DU PERSONNAGE
// ==========================================

func skill(c *Character) {
	if c.Classe == "Samurai" {
		c.Sorts = []string{"Tempete du ninja", "Coup de Poing"}
		return
	} else if c.Classe == "Cowboy" {
		c.Sorts = []string{"Slowing Time", "Coup de Poing"}
		return
	}
	c.Sorts = []string{"Coup de Bâton", "Coup de Poing"}
}

func DisplayInfo(c *Character) {
	fmt.Println("\n=== INFORMATIONS ===")
	fmt.Println("Nom :", c.Name)
	fmt.Println("Classe :", c.Classe)
	fmt.Println("Niveau :", c.Lvl)
	fmt.Printf("Expérience : %d / %d\n", c.CurrentExp, c.MaxExp)
	fmt.Println("PV :", c.CurrentHP, "/", c.MaxHP)
	fmt.Println("Mana :", c.CurrentMana, "/", c.MaxMana)
	fmt.Println("Bonus de dégâts : +", c.BonusDamage)
	fmt.Println("Or :", c.Money)

	fmt.Print("Sorts : ")
	if len(c.Sorts) == 0 {
		fmt.Print("Aucun")
	} else {
		for i, sort := range c.Sorts {
			fmt.Print(sort)
			if i < len(c.Sorts)-1 {
				fmt.Print(", ")
			}
		}
	}
	fmt.Println()

	var equipList []string
	if c.Equipment.Tete != "" {
		equipList = append(equipList, "Tête: "+c.Equipment.Tete)
	}
	if c.Equipment.Torse != "" {
		equipList = append(equipList, "Torse: "+c.Equipment.Torse)
	}
	if c.Equipment.Pieds != "" {
		equipList = append(equipList, "Pieds: "+c.Equipment.Pieds)
	}

	fmt.Print("Équipement : ")
	if len(equipList) == 0 {
		fmt.Print("Aucun")
	} else {
		for i, eq := range equipList {
			fmt.Print(eq)
			if i < len(equipList)-1 {
				fmt.Print(", ")
			}
		}
	}
	fmt.Println()

	fmt.Println("\nTapez 0 pour revenir au menu principal")
	var back int
	fmt.Scan(&back)
}

func isDead(c *Character) {
	if c.CurrentHP <= 0 {
		fmt.Println("Il a speedrun le respawn 💀")
	}
}
