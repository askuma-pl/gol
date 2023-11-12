package main

import (
	"image/color"
	"math/rand"
	"time"
)

type Cell struct {
	X          int
	Y          int
	Live       bool
	Color      color.RGBA
	LifeCycles int
}

func NewCell(x, y int) Cell {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return Cell{
		X:          x,
		Y:          y,
		Live:       RandBool(),
		Color:      color.RGBA{0, 255, 0, 255},
		LifeCycles: r.Intn(128),
	}
}

func (c *Cell) deduceLifeCycles() {
	if c.Live {
		c.LifeCycles--
		if c.LifeCycles == 0 {
			c.Live = false
		}

		if c.LifeCycles < 0 {
			c.LifeCycles = 255
		}
	}
}

func (c *Cell) resetLifeCycles() {
	c.LifeCycles = 255
}
