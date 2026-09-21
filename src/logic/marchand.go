// tâche 7 : Marchand
package main

import "fmt"

func addInventory(c *Character, item string){
	c.Inventory = append(c.Inventory, item)
}

func Marchand(c *Character){
	choix := 0
	fmt.Println("1. Potion de vie", "2. Potion de poison", "3. Quitter")
	fmt.Scanln(&choix)
	if choix == 1{
		addInventory(c, "Potion de vie")
		fmt.Println("Potion de vie a bien été ajoutée.")
	}
	if choix == 2{
		addInventory(c, "Potion de poison")
		fmt.Println("Potion de poison ajoutée...")
	}
	if choix ==3{
		fmt.Println("Sayounara, khoya.")
	}
}

func removeInventory(c *Character, item string){
	index := -1
	for i, obj := range c.Inventory {
		if obj == item {
			index = i
		}
	}
	if index == -1 {
		fmt.Println("Item introuvable dans l'inventaire.")
		return
	}
	c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)
}