package main

import "fmt"

func isDead(c *Character){
	if c.CurrentHP <= 0 {
		fmt.Println("Il a speedrun le respawn 💀")
		c.CurrentHP = c.MaxHP / 2
	}
}
