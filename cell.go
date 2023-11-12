package main

import "image/color"

type Cell struct {
	X          int
	Y          int
	Live       bool
	Color      color.RGBA
	LifeCycles int
}

func NewCell(x, y int) Cell {
	return Cell{
		X:          x,
		Y:          y,
		Live:       RandBool(),
		Color:      color.RGBA{0, 255, 0, 255},
		LifeCycles: 255,
	}
}
