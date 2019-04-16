package main

import (
	"fmt"

	"github.com/gen2brain/raylib-go/raygui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenW = 1920
	screenH = 1080
	tileW   = 128
	tileH   = 128
	w       = (screenW / tileW) //15
	h       = (screenH / tileH) //8
)

var tiles []rl.Rectangle
var show_grid = true

func PlaceTile(x, y int32, out []rl.Rectangle) []rl.Rectangle {
	newT := rl.NewRectangle(float32(x), float32(y), tileW, tileH)
	out = append(out, newT)
	return out
}

func HandleInput() {
	if rl.IsKeyDown(rl.KeyU) {
		Undo(tiles)
	}
}

func main() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(screenW, screenH, "integra -editor")

	rl.SetTargetFPS(60)
	rl.ToggleFullscreen()
	tiles = make([]rl.Rectangle, 0)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		raygui.LoadGuiStyle("gui_styles/dark.style")

		rl.DrawText("EDITOR WINDOW", screenW/2, 5, 36, rl.Black)
		if show_grid {
			for id := 0; id < w*h; id++ {
				x := id % w
				y := id / w

				rl.DrawRectangleLines(int32(x*tileW), int32(y*tileH), tileW, tileH, rl.Gray)

			}
		}

		rl.ClearBackground(rl.RayWhite)

		CreateAndDisplayEditorPanel()

		if rl.IsKeyDown(rl.KeyP) {

			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				tiles = PlaceTile(rl.GetMouseX(), rl.GetMouseY(), tiles)
			}
		}
		for _, tile := range tiles {
			rl.DrawRectangleRec(tile, rl.DarkBlue)
		}

		if rl.IsKeyPressed(rl.KeyA) {
			fmt.Printf("Tiles stored: %v\n", tiles)
		}

		HandleInput()
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
