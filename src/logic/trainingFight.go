//tâche 22
package main

import "fmt"

func trainingFight(player *Character){
	var goblin Monster
	initgoblin(&goblin)
	turn := 1
	for player.CurrentHP > 0 && goblin.CurrentHP > 0 {
		characterTurn(player, goblin)
		if goblin.CurrentHP > 0 {
			goblinPattern(goblin, player, turn)
		}
		turn++
	}
}