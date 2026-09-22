//tâche 22
package main

import "fmt"

func trainingFight(player *Character) {
	var goblin Monster
	initGoblin(&goblin)
	turn := 1
	for player.CurrentHP > 0 && goblin.CurrentHP > 0 {
		fmt.Println("Tour", turn)
		characterTurn(player, &goblin)
		if goblin.CurrentHP > 0 {
			goblinPattern(&goblin, player, turn)
		}
		turn++
	}
	if player.CurrentHP <= 0 {
		fmt.Println("Même en entraînement, tu te rates.")
	} else {
		fmt.Println("Victoire. C'est rarissime !")
	}
}