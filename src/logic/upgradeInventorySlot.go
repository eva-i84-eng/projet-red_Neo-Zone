// tâche 18
package main

import "fmt"

func upgradeInventorySlot(c *Character){
	if c.InventoryUpgrades < 3{
		c.MaxInventory += 10
		c.InventoryUpgrades += 1
	} else {
		fmt.Println("Inventaire au max. Arrête de cliquer.")
	}
}
