package main

import (
	"fmt"
	"os"

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

var tiles []tile
var show_grid = true

var lastPlaced tile

var icon *rl.Image = rl.LoadImage("icon.png")

func PlaceTile(x, y int32, out []tile) []tile {
	newTRect := rl.NewRectangle(float32(x-x%tileW), float32(y-y%tileH), tileW, tileH)
	newT := NewTile(newTRect, 0, rl.Red)
	lastPlaced = *newT
	fmt.Printf("Last place tile (%v, %v)\n", lastPlaced.graphic.X, lastPlaced.graphic.Y)
	out = append(out, *newT)
	return out
}

func SaveMap() {
	print("Saving...\n")

	write, err := os.Create("tilemap.tmdata")

	if err != nil {
		panic("Couldn't create file")
	}
	defer write.Close()

	for i := 0; i < len(tiles); i++ {
		if tiles[i].graphic != rl.NewRectangle(0, 0, 0, 0) {
			fmt.Fprintf(write, "%v, %v, %v, %v, \n",
				tiles[i].graphic.X, tiles[i].graphic.Y, tiles[i].graphic.Width, tiles[i].graphic.Height)
		}
	}
	print("Saved\n")
}

func main() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(screenW, screenH, "integra -editor")
	rl.SetWindowIcon(*icon)
	rl.SetTargetFPS(60)
	//rl.ToggleFullscreen()
	tiles = make([]tile, 0)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		raygui.LoadGuiStyle("gui_styles/dark.style")

		rl.DrawText("EDITOR WINDOW", screenW/3, 5, 36, rl.Black)
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
			rl.DrawRectangleRec(tile.graphic, tile.color)
		}

		if rl.IsKeyPressed(rl.KeyA) {
			rl.DrawText("Saved!", screenW/2, screenH/2, 30, rl.Red)
			SaveMap()
		}
		Undo()
		Clear()
		/* if rl.IsKeyPressed(rl.KeyL) {
			fmt.Println(LoadTilemapData("tilemap.imd"))
		 }*/

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
