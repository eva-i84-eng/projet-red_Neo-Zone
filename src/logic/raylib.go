package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	// Initialisation du personnage (défini dans character.go)
	noah := initCharacter("Noah", "vagabond", 1, 100, 80, []string{"Potion de mana"}, []string{"Coup de Poing"}, 100, 100)

	rl.InitWindow(800, 600, "Mon RPG en Raylib")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// Affichage des informations du personnage
		rl.DrawText("Nom: "+noah.Name, 50, 50, 20, rl.Black)
		rl.DrawText("Classe: "+noah.Classe, 50, 80, 20, rl.Black)

		rl.EndDrawing()
	}
}