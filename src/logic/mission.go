package main

func combat(perso *Character, goblin *Monster) {
	if perso.Initiative > goblin.Initiative {
		characterTurn(perso)
	} else {
		characterTurn(goblin)
	}
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
