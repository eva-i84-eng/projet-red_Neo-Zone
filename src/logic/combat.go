package main

import (
	"fmt"
	"time"
)

// ==========================================
// MÉCANIQUES DE COMBAT & EFFETS
// ==========================================

func poisonPot(c *Character) {
	for i := 0; i <= 2; i++ {
		c.CurrentHP -= 10
		fmt.Println("PV :", c.CurrentHP, "/", c.MaxHP)
		time.Sleep(1 * time.Second)
	}
}

func characterTurn(c *Character, m *Monster) bool {
	fmt.Printf("\n--- TOUR DE COMBAT ---\n")
	fmt.Printf("%s : %d/%d PV | %d/%d Mana\n", c.Name, c.CurrentHP, c.MaxHP, c.CurrentMana, c.MaxMana)
	fmt.Printf("%s : %d/%d PV\n", m.Name, m.CurrentHP, m.MaxHP)
	fmt.Println("1 - Utiliser une compétence/sort")
	fmt.Println("2 - Inventaire")
	fmt.Println("3 - Fuir le combat")

	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil {
		var dump string
		fmt.Scanln(&dump)
		fmt.Println("Choix invalide ! Le gobelin profite de votre hésitation pour attaquer !")
		return false
	}

	switch choice {
	case 1:
		fmt.Println("\n--- CHOISISSEZ UNE ATTAQUE ---")
		for i, sort := range c.Sorts {
			cost := 0
			if sort == "Coup de Poing" {
				cost = 10
			} else if sort == "Boule de Feu" {
				cost = 20
			} else if sort == "Coup de Bâton" {
				cost = 0
			}
			fmt.Printf("%d - %s (Coût: %d Mana)\n", i+1, sort, cost)
		}

		var spellChoice int
		_, errSpell := fmt.Scan(&spellChoice)
		if errSpell != nil {
			var dump string
			fmt.Scanln(&dump)
			fmt.Println("Choix d'attaque invalide ! Le gobelin en profite pour attaquer !")
			return false
		}

		if spellChoice > 0 && spellChoice <= len(c.Sorts) {
			selectedSpell := c.Sorts[spellChoice-1]
			baseDamage := 0

			if selectedSpell == "Coup de Bâton" {
				baseDamage = 8
			} else if selectedSpell == "Coup de Poing" {
				if c.CurrentMana < 10 {
					fmt.Println("Mana insuffisant ! Le gobelin en profite pour attaquer !")
					return false
				}
				c.CurrentMana -= 10
				baseDamage = 12

			} else if selectedSpell == "Boule de Feu" {
				if c.CurrentMana < 20 {
					fmt.Println("Mana insuffisant ! Le gobelin en profite pour attaquer !")
					return false
				}
				c.CurrentMana -= 20
				baseDamage = 18

			} else if selectedSpell == "Tempete du ninja" || selectedSpell == "Slowing Time" {
				baseDamage = 8
			}

			totalDamage := baseDamage + c.BonusDamage
			m.CurrentHP -= totalDamage
			fmt.Printf("Vous utilisez %s et infligez %d dégâts ! (dont +%d bonus)\n", selectedSpell, totalDamage, c.BonusDamage)

		} else {
			fmt.Println("Choix d'attaque invalide ! Le gobelin en profite pour attaquer !")
			return false
		}

		if m.CurrentHP <= 0 {
			m.CurrentHP = 0
			fmt.Println("Vous avez vaincu le monstre !")
		}
		return true

	case 2:
		AccessInventory(c)
		return true

	case 3:
		fmt.Println("Tu es une poule mouillée !")
		m.CurrentHP = -1
		return true

	default:
		fmt.Println("Choix invalide ! Le gobelin en profite pour attaquer !")
		return false
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

		if goblin.CurrentHP == -1 {
			return
		}

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
		experience(player, &goblin)
	}
}

func experience(perso *Character, goblin *Monster) {
	if goblin.CurrentHP <= 0 {
		perso.CurrentExp += goblin.Exp
		fmt.Printf("Vous gagnez %d points d'expérience !\n", goblin.Exp)

		goldEarned := 15 + (goblin.Lvl-1)*5
		perso.Money += goldEarned
		fmt.Printf("Vous ramassez %d pièces d'or ! (Total : %d po)\n", goldEarned, perso.Money)

		for perso.CurrentExp >= perso.MaxExp {
			perso.CurrentExp -= perso.MaxExp
			perso.Lvl++
			perso.MaxExp += 50
			perso.MaxHP += 15
			perso.CurrentHP = perso.MaxHP
			perso.MaxMana += 10
			perso.CurrentMana = perso.MaxMana
			perso.BonusDamage += 3

			fmt.Printf("\n🎉 NIVEAU SUPÉRIEUR ! Vous êtes maintenant niveau %d !\n", perso.Lvl)
			fmt.Printf("PV Max: %d | Mana Max: %d | Bonus Dégâts: +%d\n", perso.MaxHP, perso.MaxMana, perso.BonusDamage)
		}
	}
}