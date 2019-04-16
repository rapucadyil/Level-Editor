package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenW   = 1920
	screenH   = 1080
	maxScenes = 3
	tileW     = 128
	tileH     = 128
	w         = (screenW / tileW) //15
	h         = (screenH / tileH) //8
)

func PlaceTile(x, y int32, out *[]rl.Rectangle) []rl.Rectangle {
	newT := rl.NewRectangle(float32(x), float32(y), tileW, tileH)
	*out = append(*out, newT)
	return *out
}

func main() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(screenW, screenH, "integra -editor")

	rl.SetTargetFPS(60)

	var tiles []rl.Rectangle
	tiles = make([]rl.Rectangle, 0)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.DrawText("EDITOR WINDOW", screenW/2, 5, 36, rl.Black)
		for id := 0; id < w*h; id++ {
			x := id % w
			y := id / w

			rl.DrawRectangleLines(int32(x*tileW), int32(y*tileH), tileW, tileH, rl.Gray)

		}

		rl.ClearBackground(rl.RayWhite)

		fmt.Printf("Mouse Pos (%v, %v)\n", rl.GetMouseX(), rl.GetMouseY())

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
