package main

import "github.com/hajimehoshi/ebiten/v2"

func drawRectangle(screen *ebiten.Image, c Cell) {
	screen.Set(c.X, c.Y, c.Color)
}
