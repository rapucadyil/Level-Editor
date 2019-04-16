package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func PlaceTile(x, y int32, out *[]rl.Rectangle) []rl.Rectangle {
	newT := rl.NewRectangle(float32(x), float32(y), 128, 128)
	*out = append(*out, newT)
	return *out
}

func main() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(1280, 720, "integra -editor")

	rl.SetTargetFPS(60)

	var tiles []rl.Rectangle
	tiles = make([]rl.Rectangle, 0)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("EDITOR WINDOW", 1280/3, 5, 36, rl.LightGray)

		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			tiles = PlaceTile(rl.GetMouseX(), rl.GetMouseY(), &tiles)
		}

		for _, tile := range tiles {
			rl.DrawRectangleRec(tile, rl.DarkBlue)
		}

		if rl.IsKeyPressed(rl.KeyA) {
			fmt.Printf("Tiles stored: %v\n", tiles)
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
