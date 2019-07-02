package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	GROUND_TILE  = int32(0)
	WALL_TILE    = int32(1)
	INTERACTABLE = int32(2)
	PLAYER_START = int32(3)
)

type tile struct {
	graphic  rl.Rectangle
	tileType int32
	color    rl.Color
}

func NewTile(gr rl.Rectangle, tiletype int32, col rl.Color) *tile {
	r := new(tile)
	r.color = col
	r.graphic = gr
	r.tileType = tiletype
	return r
}
