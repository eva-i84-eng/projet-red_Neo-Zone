package main

func combat(perso *Character, goblin *Monster) {
	turn := 1
	for perso.CurrentHP > 0 && goblin.CurrentHP > 0 {
		if perso.Initiative >= goblin.Initiative {
			characterTurn(perso, goblin)
			if goblin.CurrentHP > 0 {
				goblinPattern(goblin, perso, turn)
			}
		} else {
			goblinPattern(goblin, perso, turn)
			if perso.CurrentHP > 0 {
				characterTurn(perso, goblin)
			}
		}
		isDead(perso)
		turn++
	}
	experience(perso, goblin)
}

func experience(perso *Character, goblin *Monster) {
	if goblin.CurrentHP <= 0 {
		perso.CurrentExp += goblin.Exp

		for perso.CurrentExp >= perso.MaxExp {
			perso.CurrentExp -= perso.MaxExp
			perso.Lvl++
			perso.MaxExp += 50

			perso.MaxHP += 10
			perso.CurrentHP += 10
		}
	}
}