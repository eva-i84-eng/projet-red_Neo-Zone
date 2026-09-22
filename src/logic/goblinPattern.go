//tâche 20, 22
package main

import "fmt"

func goblinPattern(c *Monster, player *Character, turn int){
	damage := c.Damage
	if turn%3 == 0{
		damage = c.Damage*2
	}
	player.CurrentHP -= damage 
	fmt.Println("Gobelin d'entrainement inflige à", player.Name, damage, "de dégâts")
	fmt.Println("PV :", player.CurrentHP, "/", player.MaxHP)
}