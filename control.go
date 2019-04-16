package main

import (
	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type EditorPanel struct {
	xpos, ypos       int32
	panel_title      string
	panel_w, panel_h int32
}

func Undo(tiles []rl.Rectangle) {
	for i := len(tiles) - 1; i >= 0; i-- {
		tiles[i] = rl.Rectangle{}
	}
}

func NewEditorPanel(title string, w, h, x, y int32) *EditorPanel {
	r := new(EditorPanel)
	r.panel_w, r.panel_h, r.xpos, r.ypos = w, h, x, y
	r.panel_title = title
	return r
}

func CreateAndDisplayEditorPanel() {
	panel := NewEditorPanel("Editor", 128, 128, 10, 10)
	bounds := rl.NewRectangle(float32(panel.xpos), float32(panel.ypos), float32(panel.panel_w), float32(panel.panel_h))
	bounds_button := rl.NewRectangle(float32(panel.xpos), float32(panel.ypos+128), float32(panel.panel_w+256), float32(panel.panel_h))
	clicked := raygui.Button(bounds_button, "Toggle Grid")

	if clicked {
		show_grid = !show_grid
	}

	raygui.Label(bounds, panel.panel_title)
}