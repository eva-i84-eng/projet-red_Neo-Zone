package main

import (
	"fmt"
	"time"
)

// ==========================================
// 1. STRUCTURES
// ==========================================

type Equipment struct {
	Tete  string
	Torse string
	Pieds string
}

type Character struct {
	Name              string
	Classe            string
	Lvl               int
	MaxHP             int
	CurrentHP         int
	MaxHPFinal        int
	Inventory         []string
	MaxInventory      int
	InventoryUpgrades int
	Money             int
	Sorts             []string
	Equipment         Equipment
	MaxMana           int
	CurrentMana       int
	Initiative        int
	CurrentExp        int
	MaxExp            int
	BonusDamage       int // Dégâts bonus gagnés avec les niveaux
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
// 2. INITIALISATIONS & MAIN
// ==========================================

func processName(input string) (string, bool) {
	var cleaned []rune

	for _, r := range input {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= 'à' && r <= 'ÿ') || (r >= 'À' && r <= 'ß') {
			cleaned = append(cleaned, r)
		}
	}
	if len(cleaned) == 0 {
		return "", false
	}
	for i := 0; i < len(cleaned); i++ {
		r := cleaned[i]
		if i == 0 {
			// Convertit la première lettre en majuscule si minuscule
			if r >= 'a' && r <= 'z' {
				cleaned[i] = r - ('a' - 'A')
			}
		} else {
			// Convertit les lettres suivantes en minuscule si majuscules
			if r >= 'A' && r <= 'Z' {
				cleaned[i] = r + ('a' - 'A')
			}
		}
	}

	return string(cleaned), true
}

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

func main() {
	var rawName string
	var name string
	var classeChoice int
	var classeName string

	fmt.Println("=== CRÉATION DU PERSONNAGE ===")
	for {
		fmt.Print("Entrez votre nom : ")
		fmt.Scan(&rawName)

		formatted, valid := processName(rawName)
		if !valid {
			fmt.Println("Erreur : Les numéros ne sont pas acceptés dans le pseudo !")
			continue
		}

		name = formatted
		break
	}

	fmt.Println("\nChoisissez votre classe :")
	fmt.Println("1 - Vagabond")
	fmt.Println("2 - Samurai")
	fmt.Println("3 - Cowboy")
	fmt.Print("Votre choix : ")
	fmt.Scan(&classeChoice)

	switch classeChoice {
	case 2:
		classeName = "Samurai"
	case 3:
		classeName = "Cowboy"
	default:
		classeName = "Vagabond"
	}

	player := initCharacter(name, classeName, 1, 100, 100, []string{"Potion de mana"}, []string{}, 100, 100)
	skill(&player)
	MainMenu(&player)
}

// ==========================================
// 3. MENU PRINCIPAL
// ==========================================

func MainMenu(c *Character) {
	for {
		fmt.Println("\n=== MENU PRINCIPAL ===")
		fmt.Println("1 - Afficher les infos du personnage")
		fmt.Println("2 - Afficher l'inventaire")
		fmt.Println("3 - Le marchand râleur")
		fmt.Println("4 - Le forgeron vantard")
		fmt.Println("5 - Qui sont-ils ?")
		fmt.Println("6 - Entraînement")
		fmt.Println("7 - Quitter")
		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			DisplayInfo(c)
		case 2:
			AccessInventory(c)
		case 3:
			Merchant(c)
		case 4:
			Blacksmith(c)
		case 5:
			WhoAreThey(c)
		case 6:
			trainingFight(c)
		case 7:
			fmt.Println("Allez ouste !")
			return
		default:
			fmt.Println("Choisis ce qui est proposé !")
		}
	}
}

// ==========================================
// 4. GESTION DU PERSONNAGE & INVENTAIRE
// ==========================================

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
	fmt.Println("Sorts :", c.Sorts)
	fmt.Println("Équipement :", c.Equipment)
	fmt.Println("Tapez 0 pour revenir au menu principal")
	var back int
	fmt.Scan(&back)
}

func AccessInventory(c *Character) {
	fmt.Println("\n=================================")
	fmt.Println("    INVENTAIRE DU PERSONNAGE    ")
	fmt.Println("=================================")
	if len(c.Inventory) == 0 {
		fmt.Println("Votre inventaire est vide")
		fmt.Println("Tapez 0 pour revenir au menu principal")
		var back int
		fmt.Scan(&back)
		return
	}

	for i, item := range c.Inventory {
		fmt.Println(i+1, "-", item)
	}

	fmt.Println("Choisissez un objet à utiliser (ou 0 pour quitter) :")
	var choice int
	fmt.Scan(&choice)

	if choice > 0 && choice <= len(c.Inventory) {
		item := c.Inventory[choice-1]
		useItem(c, item, choice-1)
	}
}

func addItem(c *Character, item string) bool {
	if len(c.Inventory) >= c.MaxInventory {
		fmt.Println("L'inventaire est plein ! Impossible d'ajouter :", item)
		return false
	}
	c.Inventory = append(c.Inventory, item)
	return true
}

func removeItems(c *Character, itemsNeeded []string) bool {
	tempInventory := append([]string{}, c.Inventory...)

	for _, needed := range itemsNeeded {
		found := false
		for i, item := range tempInventory {
			if item == needed {
				tempInventory = append(tempInventory[:i], tempInventory[i+1:]...)
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	c.Inventory = tempInventory
	return true
}

func upgradeInventorySlot(c *Character) {
	if c.InventoryUpgrades < 3 {
		c.MaxInventory += 10
		c.InventoryUpgrades += 1
		fmt.Println("Inventaire augmenté ! Capacité maximale :", c.MaxInventory)
	} else {
		fmt.Println("Inventaire au max. Arrête de cliquer.")
	}
}

func skill(c *Character) {
	if c.Classe == "Samurai" {
		c.Sorts = []string{"Tempete du ninja", "Coup de Poing"}
		return
	} else if c.Classe == "Cowboy" {
		c.Sorts = []string{"Slowing Time", "Coup de Poing"}
		return
	}
	c.Sorts = []string{"Coup de Poing"}
}

func spellBook(c *Character) {
	for _, j := range c.Sorts {
		if j == "Boule de Feu" {
			fmt.Println("Vous connaissez déjà ce sort !")
			return
		}
	}
	c.Sorts = append(c.Sorts, "Boule de Feu")
	fmt.Println("Vous avez appris : Boule de Feu !")
}

func addEquipment(char *Character, stuff string) {
	for i, obj := range char.Inventory {
		if obj == stuff {
			char.Inventory = append(char.Inventory[:i], char.Inventory[i+1:]...)
			break
		}
	}

	switch stuff {
	case "Chapeau de l'aventurier":
		if char.Equipment.Tete != "" {
			char.Inventory = append(char.Inventory, char.Equipment.Tete)
			char.MaxHP -= 10
		}
		char.Equipment.Tete = stuff
		char.MaxHP += 10

	case "Tunique de l'aventurier":
		if char.Equipment.Torse != "" {
			char.Inventory = append(char.Inventory, char.Equipment.Torse)
			char.MaxHP -= 25
		}
		char.Equipment.Torse = stuff
		char.MaxHP += 25

	case "Bottes de l'aventurier":
		if char.Equipment.Pieds != "" {
			char.Inventory = append(char.Inventory, char.Equipment.Pieds)
			char.MaxHP -= 15
		}
		char.Equipment.Pieds = stuff
		char.MaxHP += 15

	default:
		fmt.Println("Cet équipement n'existe pas :", stuff)
	}
}

// ==========================================
// 5. MARCHANDS & PNJS
// ==========================================

func Merchant(c *Character) {
	for {
		fmt.Println("\n=== MARCHAND RÂLEUR ===")
		fmt.Println("Or disponible :", c.Money)
		fmt.Println("1 - Potion de vie (3 pièces d'or)")
		fmt.Println("2 - Potion de poison (6 pièces d'or)")
		fmt.Println("3 - Livre de Sort : Boule de Feu (25 pièces d'or)")
		fmt.Println("4 - Fourrure de Loup (4 pièces d'or)")
		fmt.Println("5 - Peau de Troll (7 pièces d'or)")
		fmt.Println("6 - Cuir de Sanglier (3 pièces d'or)")
		fmt.Println("7 - Plume de Corbeau (1 pièce d'or)")
		fmt.Println("8 - Augmenter l'inventaire (30 pièces d'or)")
		fmt.Println("9 - Potion de mana (5 pièces d'or)")
		fmt.Println("0 - Retour au menu principal")
		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scan(&choice)

		var item string
		var price int

		switch choice {
		case 1:
			item = "Potion de vie"
			price = 3
		case 2:
			item = "Potion de poison"
			price = 6
		case 3:
			item = "Livre de Sort : Boule de Feu"
			price = 25
		case 4:
			item = "Fourrure de Loup"
			price = 4
		case 5:
			item = "Peau de Troll"
			price = 7
		case 6:
			item = "Cuir de Sanglier"
			price = 3
		case 7:
			item = "Plume de Corbeau"
			price = 1
		case 8:
			if c.Money >= 30 {
				c.Money -= 30
				upgradeInventorySlot(c)
			} else {
				fmt.Println("T'es fauché. Upgrade refusé.")
			}
			continue
		case 9:
			item = "Potion de mana"
			price = 5
		case 0:
			fmt.Println("C'est bon j'ai compris, dégages !")
			return
		default:
			fmt.Println("Choix invalide !")
			continue
		}

		if c.Money >= price {
			if addItem(c, item) {
				c.Money -= price
				fmt.Printf("Tu as acheté : %s pour %d pièces d'or.\n", item, price)
			}
		} else {
			fmt.Println("T'as pas assez de thunes !")
		}
	}
}

func Blacksmith(c *Character) {
	for {
		fmt.Println("\n=== LE FORGERON ===")
		fmt.Println("Or disponible :", c.Money)
		fmt.Println("1 - Chapeau de l'aventurier (1 Plume de Corbeau, 1 Cuir de Sanglier)")
		fmt.Println("2 - Tunique de l'aventurier (2 Fourrures de Loup, 1 Peau de Troll)")
		fmt.Println("3 - Bottes de l'aventurier  (1 Fourrure de Loup, 1 Cuir de Sanglier)")
		fmt.Println("0 - Retour au menu principal")
		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scan(&choice)

		var equipment string
		var neededItems []string
		cost := 5

		switch choice {
		case 1:
			equipment = "Chapeau de l'aventurier"
			neededItems = []string{"Plume de Corbeau", "Cuir de Sanglier"}
		case 2:
			equipment = "Tunique de l'aventurier"
			neededItems = []string{"Fourrure de Loup", "Fourrure de Loup", "Peau de Troll"}
		case 3:
			equipment = "Bottes de l'aventurier"
			neededItems = []string{"Fourrure de Loup", "Cuir de Sanglier"}
		case 0:
			return
		default:
			fmt.Println("Choix invalide !")
			continue
		}

		if c.Money < cost {
			fmt.Println("Pas assez de pièces !")
			continue
		}

		if !removeItems(c, neededItems) {
			fmt.Println("Il te manque des composants !")
			continue
		}

		c.Money -= cost
		addItem(c, equipment)
		fmt.Printf("Fabriqué : %s pour %d pièces d'or !\n", equipment, cost)
	}
}

func WhoAreThey(c *Character) {
	for {
		fmt.Println("\n=== LES ARTISTES CACHÉS ===")
		fmt.Println("1 - Steven Spielberg")
		fmt.Println("2 - ABBA")
		fmt.Println("0 - Retour")
		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			if c.HasTalkedToSpielberg {
				fmt.Println("Vous avez déjà salué S. Spielberg !")
			} else {
				fmt.Println("Vous avez salué S.Spielberg ! Il vous raconte une histoire et vous offre 10 pièces d'or !")
				c.Money += 10
				c.HasTalkedToSpielberg = true
			}
		case 2:
			if c.HasTalkedToABBA {
				fmt.Println("Vous avez déjà écouté ABBA !")
			} else {
				fmt.Println("ABBA vous chante une chanson, vos PV sont restaurés !")
				c.CurrentHP = c.MaxHP
				c.HasTalkedToABBA = true
			}
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

// ==========================================
// 6. COMBAT, EXPÉRIENCE & POTIONS
// ==========================================

func poisonPot(c *Character){
	for i:=0; i<=2; i++{
		c.CurrentHP -= 10
		fmt.Println("PV :", c.CurrentHP, "/", c.MaxHP)
		time.Sleep(1 * time.Second)
	}
}

func useItem(c *Character, item string, index int) {
	if item == "Potion de vie" {
		if c.CurrentHP >= c.MaxHP {
			fmt.Println("Vos points de vie sont déjà au maximum !")
			return
		}
		c.CurrentHP += 50
		if c.CurrentHP > c.MaxHP {
			c.CurrentHP = c.MaxHP
		}
		fmt.Printf("Vous utilisez une Potion de Vie. PV actuels : %d/%d\n", c.CurrentHP, c.MaxHP)
		c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)

	} else if item == "Potion de poison" {
		fmt.Println("Vous buvez la potion de poison... Mauvaise idée !")
		// Consomme l'objet dans l'inventaire
		c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)
		// Applique les dégâts sur la durée
		poisonPot(c)

	} else if item == "Potion de mana" {
		if c.CurrentMana >= c.MaxMana {
			fmt.Println("Votre mana est déjà au maximum !")
			return
		}
		c.CurrentMana += 30
		if c.CurrentMana > c.MaxMana {
			c.CurrentMana = c.MaxMana
		}
		fmt.Printf("Vous utilisez une Potion de Mana. Mana actuel : %d/%d\n", c.CurrentMana, c.MaxMana)
		c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)

	} else if item == "Chapeau de l'aventurier" || item == "Tunique de l'aventurier" || item == "Bottes de l'aventurier" {
		addEquipment(c, item)
		fmt.Printf("Vous avez équipé : %s !\n", item)

	} else if item == "Livre de Sort : Boule de Feu" {
		spellBook(c)
		c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)

	} else {
		fmt.Println("Cet objet ne peut pas être consommé ou équipé directement.")
	}
}

func characterTurn(c *Character, m *Monster) {
	fmt.Printf("\n--- TOUR DE COMBAT ---\n")
	fmt.Printf("%s : %d/%d PV | %d/%d Mana\n", c.Name, c.CurrentHP, c.MaxHP, c.CurrentMana, c.MaxMana)
	fmt.Printf("%s : %d/%d PV\n", m.Name, m.CurrentHP, m.MaxHP)
	fmt.Println("1 - Utiliser une compétence/sort")
	fmt.Println("2 - Inventaire")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("\n--- CHOISISSEZ UNE ATTAQUE ---")
		for i, sort := range c.Sorts {
			cost := 0
			if sort == "Coup de Poing" {
				cost = 10
			} else if sort == "Boule de Feu" {
				cost = 20
			}
			fmt.Printf("%d - %s (Coût: %d Mana)\n", i+1, sort, cost)
		}

		var spellChoice int
		fmt.Scan(&spellChoice)

		if spellChoice > 0 && spellChoice <= len(c.Sorts) {
			selectedSpell := c.Sorts[spellChoice-1]
			baseDamage := 0

			if selectedSpell == "Coup de Poing" {
				if c.CurrentMana < 10 {
					fmt.Println("Mana insuffisant !")
					return
				}
				c.CurrentMana -= 10
				baseDamage = 12

			} else if selectedSpell == "Boule de Feu" {
				if c.CurrentMana < 20 {
					fmt.Println("Mana insuffisant !")
					return
				}
				c.CurrentMana -= 20
				baseDamage = 18

			} else if selectedSpell == "Tempete du ninja" || selectedSpell == "Slowing Time" {
				baseDamage = 8
			}

			// Prise en compte du bonus de dégâts accordé par le niveau
			totalDamage := baseDamage + c.BonusDamage
			m.CurrentHP -= totalDamage
			fmt.Printf("Vous utilisez %s et infligez %d dégâts ! (dont +%d bonus)\n", selectedSpell, totalDamage, c.BonusDamage)

		} else {
			fmt.Println("Choix d'attaque invalide.")
			return
		}

		if m.CurrentHP <= 0 {
			m.CurrentHP = 0
			fmt.Println("Vous avez vaincu le monstre !")
		}

	case 2:
		AccessInventory(c)

	default:
		fmt.Println("Choix invalide !")
	}
}

func goblinPattern(c *Monster, player *Character, turn int) {
	damage := c.Damage
	if turn%3 == 0 {
		damage = c.Damage * 2
	}
	player.CurrentHP -= damage
	fmt.Printf("%s inflige à %s %d dégâts !\n", c.Name, player.Name, damage)
	fmt.Println("PV :", player.CurrentHP, "/", player.MaxHP)
}

func trainingFight(player *Character) {
	var goblin Monster
	initGoblin(&goblin, player.Lvl)

	fmt.Printf("\nUn %s apparaît !\n", goblin.Name)
	turn := 1

	for player.CurrentHP > 0 && goblin.CurrentHP > 0 {
		fmt.Println("\n--- Tour", turn, "---")
		characterTurn(player, &goblin)

		if goblin.CurrentHP > 0 {
			goblinPattern(&goblin, player, turn)
		}
		turn++
	}

	if player.CurrentHP <= 0 {
		fmt.Println("\nMême en entraînement, tu te rates. Tes PV sont restaurés à 1.")
		player.CurrentHP = 1
	} else {
		fmt.Println("\nVictoire !")
		experience(player, &goblin) // Gain expérience après victoire
	}
}

func experience(perso *Character, goblin *Monster) {
	if goblin.CurrentHP <= 0 {
		perso.CurrentExp += goblin.Exp
		fmt.Printf("Vous gagnez %d points d'expérience !\n", goblin.Exp)

		for perso.CurrentExp >= perso.MaxExp {
			perso.CurrentExp -= perso.MaxExp
			perso.Lvl++
			perso.MaxExp += 50
			perso.MaxHP += 15
			perso.CurrentHP = perso.MaxHP // Soin complet au passage de niveau
			perso.MaxMana += 10
			perso.CurrentMana = perso.MaxMana
			perso.BonusDamage += 3 // Augmentation des dégâts

			fmt.Printf("\n🎉 NIVEAU SUPÉRIEUR ! Vous êtes maintenant niveau %d !\n", perso.Lvl)
			fmt.Printf("PV Max: %d | Mana Max: %d | Bonus Dégâts: +%d\n", perso.MaxHP, perso.MaxMana, perso.BonusDamage)
		}
	}
}

func isDead(c *Character) {
	if c.CurrentHP <= 0 {
		fmt.Println("Il a speedrun le respawn 💀")
	}
}