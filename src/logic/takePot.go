//tâches 5 & 8
package main

import "fmt"

func takePot(c *Character) {
	index := -1
	for i, potion := range c.Inventory {
		if potion == "Potion de vie" {
			index = i
		}
	}
	if index == -1 {
		fmt.Println("Tes potions sont vides, ton avenir aussi.")
		return
	}
	c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)
	c.CurrentHP += 50
	if c.CurrentHP > c.MaxHP {
		c.CurrentHP = c.MaxHP
	}
	fmt.Println("PV :", c.CurrentHP, "/", c.MaxHP)
}