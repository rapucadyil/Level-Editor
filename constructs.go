package main

import rl "github.com/gen2brain/raylib-go/raylib"

type tile struct {
	graphic  *rl.Rectangle
	tileType int32
	color    rl.Color
}
