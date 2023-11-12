package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	scrWidth  = 256
	scrHeight = 256
	winWidht  = 1024
	winHeight = 1024
)

func RandBool() bool {
	return rand.Intn(2) == 1
}

func main() {
	ebiten.SetWindowSize(winWidht, winHeight)
	ebiten.SetWindowTitle("Game of Julek")
	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
